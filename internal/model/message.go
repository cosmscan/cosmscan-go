package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Message is a message in the transaction, it has a one-to-many relationship with the transaction
type Message struct {
	gorm.Model

	TransactionId uint           `json:"transactionId"`
	Seq           uint           `json:"seq"`
	RawData       datatypes.JSON `json:"rawData"` // RawData is the raw format of message, it's a json string
}

func (m *Message) Create(db *gorm.DB) error {
	return db.Model(&m).Create(&m).Error
}

func (m *Message) FindAllByTxId(db *gorm.DB, txId uint) ([]*Message, error) {
	var messages []*Message
	ret := db.Model(&m).Where(Message{
		TransactionId: txId,
	}).Order("seq asc").Find(&messages)

	if ret.Error != nil {
		return nil, ret.Error
	}
	return messages, nil
}
