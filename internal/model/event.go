package model

import (
	"gorm.io/gorm"
)

// Event is emitted when the transaction is executed, it has a one-to-many relationship with the transaction
type Event struct {
	gorm.Model

	ChainId uint        `json:"chainId"`
	TxId    uint        `json:"txId"`
	Height  BlockHeight `json:"height"`
	Seq     uint32      `json:"seq"`
	Type    string      `json:"type"`
	Key     string      `json:"key"`
	Value   string      `json:"value"`
	Indexed bool        `json:"indexed"`
}

// Create a new event
func (e *Event) Create(db *gorm.DB) error {
	return db.Model(&e).Create(&e).Error
}
