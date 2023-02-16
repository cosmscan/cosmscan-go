package model

import (
	"database/sql"
	"time"
)

// AccountBalance keeps the balance of the certain account
type AccountBalance struct {
	ID         int64        `db:"id" json:"id"`
	AccountId  int64        `db:"account_id" json:"accountId"`
	Amount     uint64       `db:"amount" json:"amount"`
	CoinName   string       `db:"coin_name" json:"coinName"` // CoinName is the name of coin, such as "uatom", actually, it's the denom of coin
	InsertedAt time.Time    `db:"inserted_at" json:"insertedAt"`
	UpdatedAt  sql.NullTime `db:"updated_at" json:"updatedAt"`
}
