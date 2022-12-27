package psqldb

import (
	"context"
	"cosmscan-go/db"

	sq "github.com/Masterminds/squirrel"
)

func (p *PsqlDB) InsertChain(ctx context.Context, chain *db.Chain) (int64, error) {
	var id int64
	sql, args, err := sq.Insert("chains").
		Columns("chain_id", "chain_name").
		Values(chain.ChainId, chain.ChainName).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return 0, err
	}

	if err := p.pool.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PsqlDB) InsertBlock(ctx context.Context, block *db.Block) (int64, error) {
	var id int64
	sql, args, err := sq.Insert("blocks").
		Columns(
			"chain_id", "height", "block_hash", "prev_hash", "proposer_address", "last_commit_hash",
			"data_hash", "validators_hash", "next_validators_hash", "consensus_hash", "app_hash",
			"last_result_hash", "evidence_hash", "block_time").
		Values(block.ChainId, block.Height, block.Hash, block.ParentHash, block.ProposerAddress,
			block.LastCommitHash, block.DataHash, block.ValidatorsHash, block.NextValidatorsHash,
			block.ConsensusHash, block.AppHash, block.LastResultHash, block.EvidenceHash, block.BlockTime).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return 0, err
	}

	if err := p.pool.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PsqlDB) InsertTransaction(ctx context.Context, tx *db.Transaction) (int64, error) {
	var id int64
	sql, args, err := sq.Insert("transactions").
		Columns("chain_id", "transaction_hash", "height", "code", "code_space", "tx_data",
			"raw_log", "info", "memo", "gas_wanted", "gas_used", "issued_at").
		Values(tx.ChainId, tx.Hash, tx.Height, tx.Code, tx.CodeSpace, tx.TxData,
			tx.RawLog, tx.Info, tx.Memo, tx.GasWanted, tx.GasUsed, tx.IssuedAt).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return 0, err
	}

	if err := p.pool.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PsqlDB) InsertEvent(ctx context.Context, event *db.Event) (int64, error) {
	var id int64
	sql, args, err := sq.Insert("events").
		Columns("chain_id", "tx_id", "tx_type", "block_height",
			"event_seq", "event_type", "event_key", "event_value", "indexed").
		Values(event.ChainId, event.TxId, event.TxType, event.Height,
			event.Seq, event.Type, event.Key, event.Value, event.Indexed).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return 0, err
	}

	if err := p.pool.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PsqlDB) InsertAccount(ctx context.Context, account *db.Account) (int64, error) {
	var id int64
	sql, args, err := sq.Insert("events").
		Columns("chain_id", "address").
		Values(account.ChainId, account.Address).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return 0, err
	}

	if err := p.pool.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PsqlDB) InsertMessage(ctx context.Context, message *db.Message) (int64, error) {
	var id int64
	sql, args, err := sq.Insert("events").
		Columns("transaction_id", "message_seq", "raw_data").
		Values(message.TransactionId, message.Seq, message.RawData).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return 0, err
	}

	if err := p.pool.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
