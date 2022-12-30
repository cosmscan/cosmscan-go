package indexer

import (
	"context"
	"cosmscan-go/client"
	"cosmscan-go/db"
	"cosmscan-go/db/psqldb"
	"cosmscan-go/indexer/fetcher"
	"cosmscan-go/indexer/schema"
	"errors"
	"sync"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type Indexer struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	log        *zap.SugaredLogger
	cfg        *Config
	cli        *client.Client
	storage    db.DB
}

func NewIndexer(cfg *Config) (*Indexer, error) {
	storage, err := psqldb.NewPsqlDB(&psqldb.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Database: cfg.DB.Database,
	})
	if err != nil {
		return nil, err
	}

	cli, err := client.NewClient(&client.Config{
		RPCEndpoint: cfg.RPCEndpoint,
	})
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Indexer{
		ctx:        ctx,
		cancelFunc: cancel,
		log:        zap.S(),
		cfg:        cfg,
		cli:        cli,
		storage:    storage,
	}, nil
}

func (i *Indexer) Run() error {
	wg := sync.WaitGroup{}
	i.log.Info("preparing to start indexing")

	current, err := i.pickCurrentBlock()
	if err != nil {
		return err
	}

	blockFetcher := fetcher.NewBlockFetcher(i.cli, i.storage, current)
	blockCh, err := blockFetcher.Subscribe()
	if err != nil {
		return err
	}

	accountFetcher := fetcher.NewAccountBalanceFetcher(i.cli)
	accReqCh, accResCh, err := accountFetcher.Subscribe()
	if err != nil {
		return err
	}

	// run fetcher
	wg.Add(1)
	go func() {
		defer i.Close()
		defer wg.Done()

		i.log.Info("started block fetcher")
		if err := blockFetcher.Run(); err != nil {
			i.log.Fatalw("failed to run block fetcher", "err", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer i.Close()
		defer wg.Done()

		i.log.Info("started account fetcher")
		accountFetcher.Run()
	}()

	wg.Add(1)
	go func() {
		defer i.Close()
		defer wg.Done()

		i.log.Info("started committer")
		i.startCommitter(i.ctx, blockCh, accReqCh, accResCh)
	}()

	select {
	case <-i.ctx.Done():
		i.log.Info("indexer is stopped")
		blockFetcher.Close()
	}

	wg.Wait()
	return nil
}

func (i *Indexer) Close() {
	i.cancelFunc()
}

func (i *Indexer) pickCurrentBlock() (db.BlockHeight, error) {
	block, err := i.storage.LatestBlock(context.Background())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return db.BlockHeight(i.cfg.StartBlock), nil
		}
		return 0, err
	}
	return block.Height + 1, nil
}

func (i *Indexer) startCommitter(ctx context.Context, blockCh <-chan *fetcher.FetchedBlock, accReqCh chan<- *schema.Account, accResCh <-chan *schema.AccountBalance) {
	commitCh := make(chan *schema.FullBlock, 10)
	accCommitCh := make(chan *schema.AccountBalance, 10)

	// run committer
	committer := NewCommitter(i.storage)
	go committer.Run(commitCh, accCommitCh)

	for {
		select {
		case <-ctx.Done():
			committer.Close()
			return
		case fetchedBlock := <-blockCh:
			var abciTx []*coretypes.ResultTx
			var cosmTx []*txtypes.GetTxResponse

			for _, tx := range fetchedBlock.Txs {
				abciTx = append(abciTx, tx.ABCIQueryResult)
				cosmTx = append(cosmTx, tx.CosmosQueryResult)
			}

			fullBlock, err := schema.NewFullBlock(fetchedBlock.Block, abciTx, cosmTx)
			if err != nil {
				i.log.Panicw("unexpectedly failed to create full block", "err", err)
			}

			commitCh <- fullBlock

			// extract account from full block
			accounts := schema.AccountsFromFullBlock(fullBlock)
			for _, acc := range accounts {
				accReqCh <- acc
			}
		case acc := <-accResCh:
			accCommitCh <- acc
		}
	}
}
