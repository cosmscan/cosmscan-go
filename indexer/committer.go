package indexer

import (
	"context"
	"cosmscan-go/db"
	"cosmscan-go/indexer/schema"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Committer struct {
	log        *zap.SugaredLogger
	ctx        context.Context
	cancelFunc context.CancelFunc
	storage    db.DB
}

func NewCommitter(storage db.DB) *Committer {
	ctx, cancel := context.WithCancel(context.Background())

	return &Committer{
		log:        zap.S().Named("committer"),
		ctx:        ctx,
		cancelFunc: cancel,
		storage:    storage,
	}
}

func (c *Committer) Run(blockCh chan *schema.FullBlock, accountCh chan *schema.AccountBalance) {
	type blockStat struct {
		proceeded int
		start     db.BlockHeight
		end       db.BlockHeight
	}
	type accountStat struct {
		proceeded int
	}

	stats := &struct {
		block   *blockStat
		account *accountStat
	}{
		block:   &blockStat{0, 0, 0},
		account: &accountStat{0},
	}

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			c.log.Info("committer is shutting down")
			return
		case tick := <-ticker.C:
			c.log.Infow("Number of data has been committed",
				"at", tick,
				"blocks", stats.block.proceeded,
				"block_start", stats.block.start,
				"block_end", stats.block.end,
				"accounts", stats.account.proceeded)
			stats.block.proceeded = 0
			stats.account.proceeded = 0
		case block := <-blockCh:
			if err := c.commitBlock(block); err != nil {
				// sometimes database is temporarily unavailable, in the future, we need to retry
				c.log.Fatalw("failed to commit block, this is unexpected behavior", "err", err)
			}
			c.log.Debugw("new block committed", "height", block.Block.Height)

			if stats.block.proceeded == 0 {
				stats.block.start = block.Block.Height
			} else {
				stats.block.end = block.Block.Height
			}
			stats.block.proceeded++
		case account := <-accountCh:
			accountId, err := c.storage.InsertAccount(c.ctx, &db.Account{
				ChainId: 1,
				Address: account.Account.Address,
			})
			if err != nil {
				c.log.Warn("inserting account failed with", "err", err, "account", account.Account.Address)
				continue
			}

			for _, b := range account.Balance {
				amount := b.Amount.Uint64()
				if err := c.storage.UpdateAccountBalance(c.ctx, accountId, b.Denom, amount); err != nil {
					c.log.Warn("updating account balance failed with",
						"err", err,
						"accountId", accountId,
						"denom", b.Denom,
						"amount", amount)
					continue
				}
			}

			stats.account.proceeded++
		}
	}
}

func (c *Committer) Close() {
	c.cancelFunc()
}

func (c *Committer) commitBlock(fullBlock *schema.FullBlock) error {
	if err := c.storage.WithTransaction(c.ctx, func(dbTx db.DB) error {
		fullBlock.Block.ChainId = 1
		if _, err := dbTx.InsertBlock(c.ctx, fullBlock.Block); err != nil {
			return fmt.Errorf("err insert block: %w", err)
		}

		for _, transaction := range fullBlock.Txs {
			transaction.Tx.ChainId = 1
			txId, err := dbTx.InsertTransaction(c.ctx, transaction.Tx)
			if err != nil {
				return fmt.Errorf("err insert transaction: %w", err)
			}

			for _, evt := range transaction.Events {
				evt.ChainId = 1
				evt.TxId = int(txId)
				if _, err := dbTx.InsertEvent(c.ctx, evt); err != nil {
					return fmt.Errorf("err insert event: %w", err)
				}
			}

			for _, msg := range transaction.Messages {
				msg.TransactionId = int(txId)
				if _, err := dbTx.InsertMessage(c.ctx, msg); err != nil {
					return fmt.Errorf("err insert message: %w", err)
				}
			}
		}

		return nil
	}); err != nil {
		return fmt.Errorf("err with transaction: %w", err)
	}

	return nil
}
