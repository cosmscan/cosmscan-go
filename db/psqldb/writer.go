package psqldb

import (
	"context"
	"cosmscan-go/db"
	"time"
)

func (p *PsqlDB) InsertChain(ctx context.Context, chain *db.Chain) (int64, error) {
	var id int64
	sql, args, err := psql.Insert("chains").
		Columns("chain_id", "chain_name", "inserted_at").
		Values(chain.ChainId, chain.ChainName, time.Now()).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return 0, err
	}

	if err := p.tx.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PsqlDB) InsertBlock(ctx context.Context, block *db.Block) (int64, error) {
	var id int64
	sql, args, err := psql.Insert("blocks").
		Columns(
			"chain_id", "height", "block_hash", "prev_hash", "proposer_address", "last_commit_hash",
			"data_hash", "validators_hash", "next_validators_hash", "consensus_hash", "app_hash",
			"last_result_hash", "evidence_hash", "block_time", "inserted_at").
		Values(block.ChainId, block.Height, block.Hash, block.ParentHash, block.ProposerAddress,
			block.LastCommitHash, block.DataHash, block.ValidatorsHash, block.NextValidatorsHash,
			block.ConsensusHash, block.AppHash, block.LastResultHash, block.EvidenceHash, block.BlockTime, time.Now()).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return 0, err
	}

	if err := p.tx.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PsqlDB) InsertTransaction(ctx context.Context, tx *db.Transaction) (int64, error) {
	var id int64
	sql, args, err := psql.Insert("transactions").
		Columns("chain_id", "seq", "transaction_hash", "height", "code", "code_space", "tx_data",
			"raw_log", "info", "memo", "gas_wanted", "gas_used", "issued_at", "inserted_at").
		Values(tx.ChainId, tx.Seq, tx.Hash, tx.Height, tx.Code, tx.CodeSpace, tx.TxData,
			tx.RawLog, tx.Info, tx.Memo, tx.GasWanted, tx.GasUsed, tx.IssuedAt, time.Now()).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return 0, err
	}

	if err := p.tx.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PsqlDB) InsertEvent(ctx context.Context, event *db.Event) (int64, error) {
	var id int64
	sql, args, err := psql.Insert("events").
		Columns("chain_id", "tx_id", "block_height",
			"event_seq", "event_type", "event_key", "event_value", "indexed", "inserted_at").
		Values(event.ChainId, event.TxId, event.Height,
			event.Seq, event.Type, event.Key, event.Value, event.Indexed, time.Now()).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return 0, err
	}

	if err := p.tx.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PsqlDB) InsertAccount(ctx context.Context, account *db.Account) (int64, error) {
	var id int64
	sql, args, err := psql.Insert("accounts").
		Columns("chain_id", "address", "inserted_at").
		Values(account.ChainId, account.Address, time.Now()).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return 0, err
	}

	if err := p.tx.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PsqlDB) InsertMessage(ctx context.Context, message *db.Message) (int64, error) {
	var id int64
	sql, args, err := psql.Insert("messages").
		Columns("transaction_id", "seq", "rawdata", "inserted_at").
		Values(message.TransactionId, message.Seq, message.RawData, time.Now()).
		Suffix("RETURNING \"id\"").
		ToSql()

	if err != nil {
		return 0, err
	}

	if err := p.tx.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PsqlDB) UpdateAccountBalance(ctx context.Context, accountId int64, coinName string, balance uint64) error {
	sql, args, err := psql.Select("count(*)").
		From("account_balances").
		Where("account_id = ? AND coin_name = ?", accountId, coinName).
		ToSql()
	if err != nil {
		return err
	}

	var count int
	if err := p.tx.QueryRow(ctx, sql, args...).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		sql, args, err = psql.Insert("account_balances").
			Columns("account_id", "coin_name", "amount", "inserted_at").
			Values(accountId, coinName, balance, time.Now()).
			ToSql()
	} else {
		sql, args, err = psql.Update("account_balances").
			Set("amount", balance).
			Set("updated_at", time.Now()).
			Where("account_id = ? AND coin_name = ?", accountId, coinName).
			ToSql()
	}

	if err != nil {
		return err
	}

	if _, err := p.tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}
