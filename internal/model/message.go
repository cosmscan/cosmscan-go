package model

import "gorm.io/gorm"

// Message is a message in the transaction, it has a one-to-many relationship with the transaction
type Message struct {
	gorm.Model

	TransactionId int    `json:"transactionId"`
	Seq           int    `json:"seq"`
	RawData       string `json:"rawData"` // RawData is the raw format of message, it's a json string
}
