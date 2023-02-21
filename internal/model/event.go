package model

import (
	"gorm.io/gorm"
)

// Event is emitted when the transaction is executed, it has a one-to-many relationship with the transaction
type Event struct {
	gorm.Model

	ChainId int64       `json:"chainId"`
	TxId    int         `json:"txId"`
	Height  BlockHeight `json:"height"`
	Seq     uint32      `json:"seq"`
	Type    string      `json:"type"`
	Key     string      `json:"key"`
	Value   string      `json:"value"`
	Indexed bool        `json:"indexed"`
}
