package model

import (
	"time"

	"gorm.io/gorm"
)

// Block represents a block in cosmos based blockchain
type Block struct {
	gorm.Model

	ChainId            int64       `json:"chainId"`
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
