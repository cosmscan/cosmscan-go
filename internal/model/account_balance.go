package model

import (
	"gorm.io/gorm"
)

// AccountBalance keeps the balance of the certain account
type AccountBalance struct {
	gorm.Model

	AccountId int64  `json:"accountId"`
	Amount    uint64 `json:"amount"`
	CoinName  string `json:"coinName"` // CoinName is the name of coin, such as "uatom", actually, it's the denom of coin
}
