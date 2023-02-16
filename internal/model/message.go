package model

import (
	"database/sql"
	"time"
)

// Message is a message in the transaction, it has a one-to-many relationship with the transaction
type Message struct {
	ID            int64        `db:"id" json:"id"`
	TransactionId int          `db:"transaction_id" json:"transactionId"`
	Seq           int          `db:"seq" json:"seq"`
	RawData       string       `db:"rawdata" json:"rawData"` // RawData is the raw format of message, it's a json string
	InsertedAt    time.Time    `db:"inserted_at" json:"insertedAt"`
	UpdatedAt     sql.NullTime `db:"updated_at" json:"updatedAt"`
}
