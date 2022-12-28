package psqldb

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrTxNotSupported = errors.New("connection pool doesn't support transactions")
)

type connPool struct {
	p *pgxpool.Pool
}

func (p *connPool) Begin(ctx context.Context) (pgx.Tx, error) {
	return nil, ErrTxNotSupported
}

func (p *connPool) Commit(ctx context.Context) error {
	return ErrTxNotSupported
}

func (p *connPool) Rollback(ctx context.Context) error {
	return ErrTxNotSupported
}

func (p *connPool) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return p.p.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

func (p *connPool) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return p.p.SendBatch(ctx, b)
}

func (p *connPool) LargeObjects() pgx.LargeObjects {
	panic("largeObject() method should not be called")
}

func (p *connPool) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	return nil, ErrTxNotSupported
}

func (p *connPool) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	return p.p.Exec(ctx, sql, arguments...)
}

func (p *connPool) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return p.p.Query(ctx, sql, args...)
}

func (p *connPool) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return p.p.QueryRow(ctx, sql, args...)
}

func (p *connPool) Conn() *pgx.Conn {
	panic("Conn() method in pool should not be called")
	return nil
}
