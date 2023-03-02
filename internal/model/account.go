package model

import (
	"errors"

	"gorm.io/gorm"
)

// Account describes wallet address in the blockchain
type Account struct {
	gorm.Model

	ChainId int64  `json:"chainId"`
	Address string `json:"address"`
}

// Create a new account
func (a *Account) Create(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		got := &Account{}
		err := got.FindBy(tx, a.Address, a.ChainId)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err == nil {
			return gorm.ErrRegistered
		}

		return tx.Model(&a).Create(&a).Error
	})
}

// FindBy address and chainId
func (a *Account) FindBy(db *gorm.DB, address string, chainId int64) error {
	return db.Model(&a).Where("address = ? and chain_id = ?", address, chainId).First(&a).Error
}

// FindByID find an account by ID
func (a *Account) FindByID(db *gorm.DB, id uint) error {
	return db.Model(&a).Where("id = ?", id).First(&a).Error
}
