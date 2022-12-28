package fetcher

import (
	"context"
	"cosmscan-go/client"
	"cosmscan-go/db"
	"errors"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"go.uber.org/zap"
)

type FetchedBlock struct {
	Block *coretypes.ResultBlock
	Txs   []*struct {
		ABCIQueryResult   *coretypes.ResultTx
		CosmosQueryResult *txtypes.GetTxResponse
	}
}

type BlockFetcher struct {
	cli        *client.Client
	storage    db.DB
	startBlock db.BlockHeight

	log        *zap.SugaredLogger
	ctx        context.Context
	cancelFunc context.CancelFunc
	init       bool
	subOnce    sync.Once
	channel    chan *FetchedBlock
}

func NewBlockFetcher(cli *client.Client, storage db.DB, startBlock db.BlockHeight) *BlockFetcher {
	ctx, cancel := context.WithCancel(context.Background())

	return &BlockFetcher{
		startBlock: startBlock,
		cli:        cli,
		storage:    storage,
		log:        zap.S().Named("block_fetcher"),
		ctx:        ctx,
		cancelFunc: cancel,
		init:       false,
	}
}

func (f *BlockFetcher) Subscribe() (<-chan *FetchedBlock, error) {
	if f.init {
		return nil, errors.New("already subscribed")
	}

	f.subOnce.Do(func() {
		f.channel = make(chan *FetchedBlock)
		f.init = true
	})

	return f.channel, nil
}

func (f *BlockFetcher) Run() error {
	return f.run()
}

func (f *BlockFetcher) Close() {
	f.cancelFunc()
}

func (f *BlockFetcher) run() error {
	current := f.startBlock
	f.log.Infow("start fetching blocks", "start_block", current)

	for {
		select {
		case <-f.ctx.Done():
			f.log.Infow("indexer is stopped at", "block", current)
			return nil
		default:
			b := backoff.NewExponentialBackOff()
			b.MaxElapsedTime = time.Minute
			b.MaxInterval = 10 * time.Second
			ticker := backoff.NewTicker(b)
			done := false

			for range ticker.C {
				block, retry, err := f.fetchBlock(current)
				if retry {
					if err != nil {
						f.log.Debugw("failed to fetch block, but will retry again", "block", current, "error", err)
					}
					continue
				}

				ticker.Stop()
				done = true
				if err != nil {
					return err
				}

				if !f.init {
					return errors.New("we fetched the block, but there is no subscribers")
				}

				f.channel <- block
			}

			if !done {
				return errors.New("maximum retry count exceeded while fetching block")
			}

			current++
		}
	}
}

func (f *BlockFetcher) fetchBlock(height db.BlockHeight) (ret *FetchedBlock, retry bool, err error) {
	var result FetchedBlock

	latestHeight, err := f.cli.LatestBlockNumber(f.ctx)
	if err != nil {
		return nil, true, err
	}

	latest := db.BlockHeight(latestHeight)
	if latest < height {
		return nil, true, nil
	}

	block, err := f.cli.Block(f.ctx, int64(height))
	if err != nil {
		return nil, false, err
	}

	result.Block = block
	result.Txs = make([]*struct {
		ABCIQueryResult   *coretypes.ResultTx
		CosmosQueryResult *txtypes.GetTxResponse
	}, 0)

	for _, tx := range block.Block.Txs {
		abciQueryResult, err := f.cli.ABCITransactionByHash(f.ctx, tx.Hash())
		if err != nil {
			return nil, false, err
		}

		cosmosQueryResult, err := client.TransactionByHash(f.ctx, abciQueryResult.Hash.String(), f.cli)
		if err != nil {
			return nil, false, err
		}

		result.Txs = append(result.Txs, &struct {
			ABCIQueryResult   *coretypes.ResultTx
			CosmosQueryResult *txtypes.GetTxResponse
		}{
			ABCIQueryResult:   abciQueryResult,
			CosmosQueryResult: cosmosQueryResult,
		})
	}

	return &result, false, nil
}
