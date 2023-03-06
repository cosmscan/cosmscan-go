package model

import (
	"gorm.io/gorm"
)

// Event is emitted when the transaction is executed, it has a one-to-many relationship with the transaction
type Event struct {
	gorm.Model

	ChainId uint   `json:"chainId"`
	TxId    uint   `json:"txId"`
	Height  uint32 `json:"height"`
	Seq     uint32 `json:"seq"`
	Type    string `json:"type"`
	Key     string `json:"key"`
	Value   string `json:"value"`
	Indexed bool   `json:"indexed"`
}

// Create a new event
func (e *Event) Create(db *gorm.DB) error {
	return db.Model(&e).Create(&e).Error
}

// FindAllByTxId find all events by tx id
func (e *Event) FindAllByTxId(db *gorm.DB, txId uint) ([]*Event, error) {
	var events []*Event
	ret := db.Model(&e).Order("seq asc").Where(Event{
		TxId: txId,
	}).Find(&events)

	if ret.Error != nil {
		return nil, ret.Error
	}

	return events, nil
}
