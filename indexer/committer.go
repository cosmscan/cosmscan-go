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

func (c *Committer) Run(blockCh chan *schema.FullBlock) {
	var cnt int
	var start db.BlockHeight
	var end db.BlockHeight

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			c.log.Info("committer is shutting down")
			return
		case tick := <-ticker.C:
			cnt = 0
			c.log.Infow("blocks has been committed", "at", tick, "cnt", cnt, "start", start, "end", end)
		case block := <-blockCh:
			if err := c.commitBlock(block); err != nil {
				// sometimes database is temporarily unavailable, in the future, we need to retry
				c.log.Fatalw("failed to commit block, this is unexpected behavior", "err", err)
			}
			c.log.Debugw("new block committed", "height", block.Block.Height)

			if cnt == 0 {
				start = block.Block.Height
			} else {
				end = block.Block.Height
			}
			cnt++
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
