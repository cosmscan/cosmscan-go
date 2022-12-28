package indexer

import (
	"context"
	"cosmscan-go/client"
	"cosmscan-go/db"
	"cosmscan-go/db/psqldb"
	"cosmscan-go/indexer/fetcher"
	"cosmscan-go/indexer/schema"
	"errors"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"sync"

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

	// run fetcher
	wg.Add(1)
	go func() {
		i.log.Info("started block fetcher")
		defer wg.Done()
		if err := blockFetcher.Run(); err != nil {
			i.log.Fatalw("failed to run block fetcher", "err", err)
		}
		i.Close()
	}()

	wg.Add(1)
	go func() {
		i.log.Info("started committer")
		defer wg.Done()
		i.startCommitter(i.ctx, blockCh)
		i.Close()
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

func (i *Indexer) startCommitter(ctx context.Context, blockCh <-chan *fetcher.FetchedBlock) {
	commitCh := make(chan *schema.FullBlock, 10)

	// run committer
	committer := NewCommitter(i.storage)
	go committer.Run(commitCh)

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
		}
	}
}
