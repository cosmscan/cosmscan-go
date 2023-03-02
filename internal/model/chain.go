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

// Create a new chain
func (c *Chain) Create(db *gorm.DB) error {
	return db.Model(&c).Create(&c).Error
}

// FindByID find a chain by id
func (c *Chain) FindByID(db *gorm.DB, chainId uint) error {
	return db.Where("id = ?", chainId).First(&c).Error
}
