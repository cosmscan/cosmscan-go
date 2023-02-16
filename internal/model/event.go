package model

import (
	"database/sql"
	"time"
)

// Event is emitted when the transaction is executed, it has a one-to-many relationship with the transaction
type Event struct {
	ID         int64        `db:"id" json:"id"`
	ChainId    int64        `db:"chain_id" json:"chainId"`
	TxId       int          `db:"tx_id" json:"txId"`
	Height     BlockHeight  `db:"block_height" json:"height"`
	Seq        uint32       `db:"event_seq" json:"seq"`
	Type       string       `db:"event_type" json:"type"`
	Key        string       `db:"event_key" json:"key"`
	Value      string       `db:"event_value" json:"value"`
	Indexed    bool         `db:"indexed" json:"indexed"`
	InsertedAt time.Time    `db:"inserted_at" json:"insertedAt"`
	UpdatedAt  sql.NullTime `db:"updated_at" json:"updatedAt"`
}
