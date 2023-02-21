package model

import (
	"gorm.io/gorm"
)

// Account describes wallet address in the blockchain
type Account struct {
	gorm.Model

	ChainId int64  `json:"chainId"`
	Address string `json:"address"`
}
