package model

import (
	"gorm.io/datatypes"
	"time"

	"gorm.io/gorm"
)

// Transaction represents a transaction in cosmos based blockchain
type Transaction struct {
	gorm.Model

	ChainId   uint           `json:"chainId"`
	Hash      string         `json:"hash"`
	Height    uint32         `json:"height"`
	Code      int            `json:"code"`
	CodeSpace string         `json:"codeSpace"`
	TxData    string         `json:"txData"`
	RawLog    datatypes.JSON `json:"rawLog"`
	Info      string         `json:"info"`
	Memo      string         `json:"memo"`
	Seq       int            `json:"seq"`
	GasWanted uint64         `json:"gasWanted"`
	GasUsed   uint64         `json:"gasUsed"`
	IssuedAt  time.Time      `json:"issuedAt"`
}

// Create a new transaction
func (t *Transaction) Create(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

// FindByHash find a transaction by hash
func (t *Transaction) FindByHash(db *gorm.DB, chainId uint, hash string) error {
	return db.Model(&t).Where(Transaction{
		ChainId: chainId,
		Hash:    hash,
	}).First(&t).Error
}

// FindAllByHeight find a transactions by height
func (t *Transaction) FindAllByHeight(db *gorm.DB, chainId uint, height uint32) ([]*Transaction, error) {
	var transactions []*Transaction
	ret := db.Model(&t).Where(Transaction{
		ChainId: chainId,
		Height:  height,
	}).Order("seq asc").Find(&transactions)

	if ret.Error != nil {
		return nil, ret.Error
	}

	return transactions, nil
}
