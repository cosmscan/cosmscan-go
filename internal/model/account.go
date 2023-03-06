package model

import (
	"gorm.io/gorm"
)

// Account describes wallet address in the blockchain
type Account struct {
	gorm.Model

	ChainId uint   `json:"chainId"`
	Address string `json:"address"`
}

// Create a new account
func (a *Account) Create(db *gorm.DB) error {
	return db.Model(&a).Create(&a).Error
}

// FindBy address and chainId
func (a *Account) FindBy(db *gorm.DB, address string, chainId uint) error {
	return db.Model(&a).Where("address = ? and chain_id = ?", address, chainId).First(&a).Error
}

// FindByID find an account by ID
func (a *Account) FindByID(db *gorm.DB, id uint) error {
	return db.Model(&a).Where("id = ?", id).First(&a).Error
}
