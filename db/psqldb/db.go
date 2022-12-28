package psqldb

import (
	"context"
	"cosmscan-go/db"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// PsqlDB means Relational Database
type PsqlDB struct {
	tx pgx.Tx
}

func NewPsqlDB(config *Config) (*PsqlDB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the postgres database: %v", err)
	}

	return &PsqlDB{
		tx: &connPool{pool},
	}, nil
}

func (p *PsqlDB) WithTransaction(ctx context.Context, fn func(tx db.DB) error) error {
	if cp, ok := p.tx.(*connPool); ok {
		ptx, err := cp.p.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			return fmt.Errorf("failed to begin transction: %v", err)
		}

		if err := fn(&PsqlDB{tx: ptx}); err != nil {
			rollbackErr := ptx.Rollback(ctx)
			if rollbackErr != nil {
				return fmt.Errorf("failed to rollback transaction: %v", rollbackErr)
			}

			return fmt.Errorf("failed to execute transaction: %v, it has been rolled back", err)
		}

		if err := ptx.Commit(ctx); err != nil {
			return fmt.Errorf("failed to commit transaction: %v", err)
		}
	} else {
		return fmt.Errorf("cannot start transaction on a transaction")
	}

	return nil
}

func (p *PsqlDB) Close() error {
	if cp, ok := p.tx.(*connPool); ok {
		cp.p.Close()
	}

	return nil
}
