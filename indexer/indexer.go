package indexer

import (
	"context"
	"cosmscan-go/client"
	"cosmscan-go/db"
	"cosmscan-go/db/psqldb"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type Indexer struct {
	log     *zap.SugaredLogger
	cfg     *Config
	cli     *client.Client
	storage db.DB
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

	return &Indexer{
		log:     zap.S(),
		cfg:     cfg,
		cli:     cli,
		storage: storage,
	}, nil
}

func (i *Indexer) Run(ctx context.Context) error {
	i.log.Info("started indexer app")
	return i.run(ctx)
}

func (i *Indexer) run(ctx context.Context) error {
	current, err := i.pickCurrentBlock()
	if err != nil {
		return err
	}
	i.log.Infow("start indexing", "start_block", current)

	for {
		select {
		case <-ctx.Done():
			i.log.Infow("indexer is stopped at", "block", current)
			return nil
		default:
			latestHeight, err := i.cli.LatestBlockNumber(ctx)
			if err != nil {
				return err
			}

			latest := db.BlockHeight(latestHeight)
			if latest < current {
				i.log.Debugw("waiting for new block", "current", current, "latest", latest)
				<-time.After(time.Second * 5)
				continue
			}

			block, err := i.cli.Block(ctx, int64(current))
			if err != nil {
				return err
			}

			i.log.Info("fetched new block", "height", current, "hash", block.Block.Hash().String())

			for _, tx := range block.Block.Txs {
				abciTx, err := i.cli.ABCITransactionByHash(ctx, tx.Hash())
				if err != nil {
					return err
				}

				i.log.Info("found new abci tx", "hash", abciTx.Hash.String())

				cosmTx, err := client.TransactionByHash(ctx, string(tx.Hash()), i.cli)
				if err != nil {
					return err
				}
				i.log.Info("found new cosmos tx", "hash", cosmTx.Tx.String(), "message", cosmTx.TxResponse.Data)
				i.log.Debug("raw cosmos messages", "message", cosmTx.TxResponse.Logs)
			}

			current++
		}
	}
}

func (i *Indexer) pickCurrentBlock() (db.BlockHeight, error) {
	block, err := i.storage.LatestBlock(context.Background())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return db.BlockHeight(i.cfg.StartBlock), nil
		}
		i.log.Errorw("failed to query latest block on the storage", "err", err)
		return 0, err
	}
	return block.Height, nil
}
