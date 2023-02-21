package model

import (
	"gorm.io/gorm"
)

// Chain represents a specific blockchain
type Chain struct {
	gorm.Model

	ChainId   string `json:"chainId"`
	ChainName string `json:"chainName"`
}
