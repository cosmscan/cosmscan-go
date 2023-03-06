package model

import (
	"gorm.io/gorm"
)

// AccountBalance keeps the balance of the certain account
type AccountBalance struct {
	gorm.Model

	AccountID uint   `json:"accountId"`
	Amount    uint64 `json:"amount"`
	CoinName  string `json:"coinName"` // CoinName is the name of coin, such as "uatom", actually, it's the denom of coin
}

// Create a new account balance
func (a *AccountBalance) Create(db *gorm.DB) error {
	return db.Model(&a).Create(&a).Error
}
