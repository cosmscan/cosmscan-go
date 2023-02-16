package model

import (
	"database/sql"
	"time"
)

// Block represents a block in cosmos based blockchain
type Block struct {
	ID                 int64        `db:"id" json:"id"`
	ChainId            int64        `db:"chain_id" json:"chainId"`
	Height             BlockHeight  `db:"height" json:"height"`
	Hash               string       `db:"block_hash" json:"hash"`
	ParentHash         string       `db:"prev_hash" json:"parentHash"`
	ProposerAddress    string       `db:"proposer_address" json:"proposerAddress"`
	LastCommitHash     string       `db:"last_commit_hash" json:"lastCommitHash"`
	DataHash           string       `db:"data_hash" json:"dataHash"`
	ValidatorsHash     string       `db:"validators_hash" json:"validatorsHash"`
	NextValidatorsHash string       `db:"next_validators_hash" json:"nextValidatorsHash"`
	ConsensusHash      string       `db:"consensus_hash" json:"consensusHash"`
	AppHash            string       `db:"app_hash" json:"appHash"`
	LastResultHash     string       `db:"last_result_hash" json:"lastResultHash"`
	EvidenceHash       string       `db:"evidence_hash" json:"evidenceHash"`
	BlockTime          time.Time    `db:"block_time" json:"blockTime"`
	InsertedAt         time.Time    `db:"inserted_at" json:"insertedAt"`
	UpdatedAt          sql.NullTime `db:"updated_at" json:"updatedAt"`
}
