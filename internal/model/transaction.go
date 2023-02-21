package model

import (
	"time"

	"gorm.io/gorm"
)

// Transaction represents a transaction in cosmos based blockchain
type Transaction struct {
	gorm.Model

	ChainId   int64       `json:"chainId"`
	Hash      string      `json:"hash"`
	Height    BlockHeight `json:"height"`
	Code      int         `json:"code"`
	CodeSpace string      `json:"codeSpace"`
	TxData    string      `json:"txData"`
	RawLog    string      `json:"rawLog"`
	Info      string      `json:"info"`
	Memo      string      `json:"memo"`
	Seq       int         `json:"seq"`
	GasWanted uint64      `json:"gasWanted"`
	GasUsed   uint64      `json:"gasUsed"`
	IssuedAt  time.Time   `json:"issuedAt"`
}
