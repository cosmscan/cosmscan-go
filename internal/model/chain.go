package model

import (
	"database/sql"
	"time"
)

// Chain represents a specific blockchain
type Chain struct {
	ID         int64          `db:"id" json:"id"`
	ChainId    string         `db:"chain_id" json:"chainId"`
	ChainName  string         `db:"chain_name" json:"chainName"`
	IconUrl    sql.NullString `db:"icon_url" json:"iconUrl"`
	Website    sql.NullString `db:"website" json:"website"`
	InsertedAt time.Time      `db:"inserted_at" json:"insertedAt"`
	UpdatedAt  sql.NullTime   `db:"updated_at" json:"updated_at"`
}
