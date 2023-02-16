package model

import (
	"database/sql"
	"time"
)

// Account describes wallet address in the blockchain
type Account struct {
	ID         int64        `db:"id" json:"id"`
	ChainId    int64        `db:"chain_id" json:"chainId"`
	Address    string       `db:"address" json:"address"`
	InsertedAt time.Time    `db:"inserted_at" json:"insertedAt"`
	UpdatedAt  sql.NullTime `db:"updated_at" json:"updatedAt"`
}
