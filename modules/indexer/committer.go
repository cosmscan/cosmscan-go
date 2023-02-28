package indexer

import (
	"context"
	"cosmscan-go/internal/db"
	schema2 "cosmscan-go/modules/indexer/schema"
	"cosmscan-go/pkg/log"
	"time"
)

type Committer struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	chainId    int64
	storage    *db.DB
}

func NewCommitter(storage *db.DB, chainId int64) *Committer {
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
		start     int64
		end       int64
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
		case <-ticker.C:
			// update metric and flush
		case <-blockCh:
			// create block
			stats.block.proceeded++
		case <-accountCh:
			// create account
			stats.account.proceeded++
		}
	}
}

func (c *Committer) Close() {
	c.cancelFunc()
}

func (c *Committer) commitBlock(fullBlock *schema2.FullBlock) error {

	return nil
}
