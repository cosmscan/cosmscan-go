package indexer

import (
	"context"
	"cosmscan-go/config"
	"cosmscan-go/db"
	"cosmscan-go/db/psqldb"
	"cosmscan-go/internal/client"
	fetcher2 "cosmscan-go/modules/indexer/fetcher"
	schema2 "cosmscan-go/modules/indexer/schema"
	"cosmscan-go/pkg/log"
	"errors"
	"sync"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/jackc/pgx/v5"
)

type Indexer struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	cfg        *config.IndexerConfig
	cli        *client.Client
	storage    db.DB
}

func NewIndexer(cfg *config.IndexerConfig) (*Indexer, error) {
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
		cfg:        cfg,
		cli:        cli,
		storage:    storage,
	}, nil
}

func (i *Indexer) Run() error {
	wg := sync.WaitGroup{}
	log.Logger.Info("preparing to start indexing")

	chainId := i.loadOrStoreChainId()

	// pick last committed block
	current, err := i.pickCurrentBlock()
	if err != nil {
		return err
	}

	blockFetcher := fetcher2.NewBlockFetcher(i.cli, i.storage, current)
	blockCh, err := blockFetcher.Subscribe()
	if err != nil {
		return err
	}

	accountFetcher := fetcher2.NewAccountBalanceFetcher(i.cli)
	accReqCh, accResCh, err := accountFetcher.Subscribe()
	if err != nil {
		return err
	}

	// run fetcher
	wg.Add(1)
	go func() {
		defer i.Close()
		defer wg.Done()

		log.Logger.Info("started block fetcher")
		if err := blockFetcher.Run(); err != nil {
			log.Logger.Fatalw("failed to run block fetcher", "err", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer i.Close()
		defer wg.Done()

		log.Logger.Info("started account fetcher")
		accountFetcher.Run()
	}()

	wg.Add(1)
	go func() {
		defer i.Close()
		defer wg.Done()

		log.Logger.Info("started committer")
		i.startCommitter(i.ctx, chainId, blockCh, accReqCh, accResCh)
	}()

	select {
	case <-i.ctx.Done():
		log.Logger.Info("indexer is stopped")
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

func (i *Indexer) startCommitter(ctx context.Context, chainId int64, blockCh <-chan *fetcher2.FetchedBlock, accReqCh chan<- *schema2.Account, accResCh <-chan *schema2.AccountBalance) {
	commitCh := make(chan *schema2.FullBlock, 10)
	accCommitCh := make(chan *schema2.AccountBalance, 10)

	// run committer
	committer := NewCommitter(i.storage, chainId)
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

			fullBlock, err := schema2.NewFullBlock(fetchedBlock.Block, abciTx, cosmTx)
			if err != nil {
				log.Logger.Panicw("unexpectedly failed to create full block", "err", err)
			}

			commitCh <- fullBlock

			// extract account from full block and send to the channel
			go func(fb *schema2.FullBlock) {
				accounts := schema2.AccountsFromFullBlock(fb)
				for _, acc := range accounts {
					accReqCh <- acc
				}
			}(fullBlock)
		case acc := <-accResCh:
			accCommitCh <- acc
		}
	}
}

// loadOrStoreChainId load or store the chain information from configs.
// it stores the new chain if it doesn't exist with given name.
func (i *Indexer) loadOrStoreChainId() int64 {
	chain, err := i.storage.FindChainByName(context.Background(), i.cfg.Chain.Name)
	if err == nil {
		return chain.ID
	}

	if err != db.ErrNotFound {
		log.Logger.Fatalw("failed to find chain with unknown reason", "err", err)
		return 0
	} else {
		insertedId, err := i.storage.InsertChain(context.Background(), &db.Chain{
			ChainId:   i.cfg.Chain.ID,
			ChainName: i.cfg.Chain.Name,
		})
		if err != nil {
			log.Logger.Fatalw("failed to create a new chain", "err", err)
		}
		return insertedId
	}
}
