package model

import (
	"time"

	"gorm.io/gorm"
)

// Block represents a block in cosmos based blockchain
type Block struct {
	gorm.Model

	ChainID            uint        `json:"chainId"`
	Height             BlockHeight `json:"height"`
	Hash               string      `json:"hash"`
	ParentHash         string      `json:"parentHash"`
	ProposerAddress    string      `json:"proposerAddress"`
	LastCommitHash     string      `json:"lastCommitHash"`
	DataHash           string      `json:"dataHash"`
	ValidatorsHash     string      `json:"validatorsHash"`
	NextValidatorsHash string      `json:"nextValidatorsHash"`
	ConsensusHash      string      `json:"consensusHash"`
	AppHash            string      `json:"appHash"`
	LastResultHash     string      `json:"lastResultHash"`
	EvidenceHash       string      `json:"evidenceHash"`
	BlockTime          time.Time   `json:"blockTime"`
}

// Create a new block
func (b *Block) Create(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		c := &Chain{}
		if err := c.FindByID(tx, b.ChainID); err != nil {
			return err
		}

		return tx.Model(&b).Create(&b).Error
	})
}

// FindByHash find a block by hash
func (b *Block) FindByHash(db *gorm.DB, hash string) error {
	return db.Model(&b).Where(Block{
		Hash: hash,
	}).First(&b).Error
}
