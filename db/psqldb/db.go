package psqldb

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
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
	pool *pgxpool.Pool
}

func NewPsqlDB(config *Config) (*PsqlDB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the postgres database: %v", err)
	}

	return &PsqlDB{
		pool: pool,
	}, nil
}

func (p *PsqlDB) Close() error {
	p.pool.Close()
	return nil
}
