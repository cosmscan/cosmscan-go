package model

import (
	"database/sql"
	"time"
)

// Transaction represents a transaction in cosmos based blockchain
type Transaction struct {
	ID         int64        `db:"id" json:"id"`
	ChainId    int64        `db:"chain_id" json:"chainId"`
	Hash       string       `db:"transaction_hash" json:"hash"`
	Height     BlockHeight  `db:"height" json:"height"`
	Code       int          `db:"code" json:"code"`
	CodeSpace  string       `db:"code_space" json:"codeSpace"`
	TxData     string       `db:"tx_data" json:"txData"`
	RawLog     string       `db:"raw_log" json:"rawLog"`
	Info       string       `db:"info" json:"info"`
	Memo       string       `db:"memo" json:"memo"`
	Seq        int          `db:"seq" json:"seq"`
	GasWanted  uint64       `db:"gas_wanted" json:"gasWanted"`
	GasUsed    uint64       `db:"gas_used" json:"gasUsed"`
	IssuedAt   time.Time    `db:"issued_at" json:"issuedAt"`
	InsertedAt time.Time    `db:"inserted_at" json:"insertedAt"`
	UpdatedAt  sql.NullTime `db:"updated_at" json:"updatedAt"`
}
