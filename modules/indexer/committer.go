package indexer

import (
	"context"
	"cosmscan-go/db"
	schema2 "cosmscan-go/modules/indexer/schema"
	"cosmscan-go/pkg/log"
	"fmt"
	"time"
)

type Committer struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	chainId    int64
	storage    db.DB
}

func NewCommitter(storage db.DB, chainId int64) *Committer {
	ctx, cancel := context.WithCancel(context.Background())

	return &Committer{
		ctx:        ctx,
		cancelFunc: cancel,
		chainId:    chainId,
		storage:    storage,
	}
}

func (c *Committer) Run(blockCh chan *schema2.FullBlock, accountCh chan *schema2.AccountBalance) {
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
			log.Logger.Info("committer is shutting down")
			return
		case tick := <-ticker.C:
			log.Logger.Infow("Number of data has been committed",
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
				log.Logger.Fatalw("failed to commit block, this is unexpected behavior", "err", err)
			}
			log.Logger.Debugw("new block committed", "height", block.Block.Height)

			if stats.block.proceeded == 0 {
				stats.block.start = block.Block.Height
			} else {
				stats.block.end = block.Block.Height
			}
			stats.block.proceeded++
		case account := <-accountCh:
			var accountId int64

			got, err := c.storage.FindAccountByAddress(c.ctx, 1, account.Account.Address)
			if err != nil && err != db.ErrNotFound {
				log.Logger.Fatalw("failed to check account existence", "err", err)
			}

			if err == db.ErrNotFound {
				accountId, err = c.storage.InsertAccount(c.ctx, &db.Account{
					ChainId: c.chainId,
					Address: account.Account.Address,
				})
				if err != nil {
					log.Logger.Warn("inserting account failed with", "err", err, "account", account.Account.Address)
					continue
				}
			} else {
				accountId = got.ID
			}

			for _, b := range account.Balance {
				amount := b.Amount.Uint64()
				if err := c.storage.UpdateAccountBalance(c.ctx, accountId, b.Denom, amount); err != nil {
					log.Logger.Warn("updating account balance failed with",
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

func (c *Committer) commitBlock(fullBlock *schema2.FullBlock) error {
	if err := c.storage.WithTransaction(c.ctx, func(dbTx db.DB) error {
		fullBlock.Block.ChainId = c.chainId
		if _, err := dbTx.InsertBlock(c.ctx, fullBlock.Block); err != nil {
			return fmt.Errorf("err insert block: %w", err)
		}

		for _, transaction := range fullBlock.Txs {
			transaction.Tx.ChainId = c.chainId
			txId, err := dbTx.InsertTransaction(c.ctx, transaction.Tx)
			if err != nil {
				return fmt.Errorf("err insert transaction: %w", err)
			}

			for _, evt := range transaction.Events {
				evt.ChainId = c.chainId
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
