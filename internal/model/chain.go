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
	return db.Where(Chain{
		Model: gorm.Model{ID: chainId},
	}).First(&c).Error
}

// FindByChainID find a chain by chainId which is unique identifier as string for a specific chain
func (c *Chain) FindByChainID(db *gorm.DB, chainId string) error {
	return db.Where(Chain{
		ChainId: chainId,
	}).First(&c).Error
}

// FindAll returns all registered chains
func (c *Chain) FindAll(db *gorm.DB) ([]Chain, error) {
	var chains []Chain
	return chains, db.Find(&chains).Error
}
