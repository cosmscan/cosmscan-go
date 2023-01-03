package psqldb

import (
	"context"
	"cosmscan-go/db"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

func (p *PsqlDB) Block(ctx context.Context, height db.BlockHeight) (*db.Block, error) {
	var block db.Block

	sql, args, err := psql.Select("*").
		From("blocks").
		Where(sq.Eq{"height": height}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row, err := p.tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	if err := pgxscan.ScanOne(&block, row); err != nil {
		return nil, err
	}

	return &block, nil
}

func (p *PsqlDB) LatestBlock(ctx context.Context) (*db.Block, error) {
	var block db.Block

	sql, args, err := psql.Select("*").
		From("blocks").
		OrderBy("height desc").
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	row, err := p.tx.Query(ctx, sql, args...)
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

	sql, args, err := psql.Select("*").
		From("blocks").
		Where(sq.Eq{"block_hash": hash}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row, err := p.tx.Query(ctx, sql, args...)
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

	sql, args, err := psql.Select("*").From("transactions").Where(sq.Eq{"transaction_hash": hash}).ToSql()
	if err != nil {
		return nil, err
	}

	row, err := p.tx.Query(ctx, sql, args...)
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

	sql, args, err := psql.Select("*").From("events").Where(sq.Eq{"tx_id": txId}).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(ctx, p.tx, &events, sql, args...); err != nil {
		return nil, err
	}

	return events, nil
}

func (p *PsqlDB) MessagesInTx(ctx context.Context, txId int) ([]*db.Message, error) {
	var messages []*db.Message

	sql, args, err := psql.Select("*").From("messages").Where(sq.Eq{"transaction_id": txId}).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(ctx, p.tx, &messages, sql, args...); err != nil {
		return nil, err
	}

	return messages, nil
}

func (p *PsqlDB) FindAccountByAddress(ctx context.Context, chainId int64, address string) (*db.Account, error) {
	var acc db.Account

	sql, args, err := psql.Select("*").
		From("accounts").
		Where(sq.Eq{"chain_id": chainId, "address": address}).ToSql()

	if err != nil {
		return nil, err
	}

	row, err := p.tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	if err := pgxscan.ScanOne(&acc, row); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, db.ErrNotFound
		} else {
			return nil, err
		}
	}

	return &acc, nil
}

func (p *PsqlDB) FindChainByName(ctx context.Context, name string) (*db.Chain, error) {
	var chain db.Chain

	sql, args, err := psql.Select("*").
		From("chains").
		Where(sq.Eq{"chain_name": name}).ToSql()

	if err != nil {
		return nil, err
	}

	row, err := p.tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	if err := pgxscan.ScanOne(&chain, row); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, db.ErrNotFound
		} else {
			return nil, err
		}
	}

	return &chain, nil
}
