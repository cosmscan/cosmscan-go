package psqldb

import (
	"context"
	"cosmscan-go/db"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (p *PsqlDB) Block(ctx context.Context, height db.BlockHeight) (*db.Block, error) {
	var block db.Block

	sql, args, err := sq.Select("*").From("blocks").Where(sq.Eq{"height": height}).ToSql()
	if err != nil {
		return nil, err
	}

	row, err := p.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	if err := pgxscan.ScanOne(&block, row); err != nil {
		return nil, err
	}

	return &block, nil
}

func (p *PsqlDB) BlockByHash(ctx context.Context, hash string) (*db.Block, error) {
	var block db.Block

	sql, args, err := sq.Select("*").From("blocks").Where(sq.Eq{"block_hash": hash}).ToSql()
	if err != nil {
		return nil, err
	}

	row, err := p.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	if err := pgxscan.ScanOne(&block, row); err != nil {
		return nil, err
	}

	return &block, nil
}

func (p *PsqlDB) Transaction(ctx context.Context, hash string) (*db.Transaction, error) {
	var tx db.Transaction

	sql, args, err := sq.Select("*").From("transactions").Where(sq.Eq{"transaction_hash": hash}).ToSql()
	if err != nil {
		return nil, err
	}

	row, err := p.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	if err := pgxscan.ScanOne(&tx, row); err != nil {
		return nil, err
	}

	return &tx, nil
}

func (p *PsqlDB) EventsInTx(ctx context.Context, txId int) ([]*db.Event, error) {
	var events []*db.Event

	sql, args, err := sq.Select("*").From("events").Where(sq.Eq{"tx_id": txId}).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(ctx, p.pool, &events, sql, args...); err != nil {
		return nil, err
	}

	return events, nil
}

func (p *PsqlDB) MessagesInTx(ctx context.Context, txId int) ([]*db.Message, error) {
	var messages []*db.Message

	sql, args, err := sq.Select("*").From("messages").Where(sq.Eq{"transaction_id": txId}).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(ctx, p.pool, &messages, sql, args...); err != nil {
		return nil, err
	}

	return messages, nil
}
