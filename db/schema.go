package db

import (
	"database/sql"
	"time"
)

type TxType int
type BlockHeight uint64

const (
	NormalTx TxType = iota
	BeginBlock
	EndBlock
)

// Chain represents a specific blockchain
type Chain struct {
	ID         int64          `db:"id" json:"id"`
	ChainId    string         `db:"chain_id" json:"chainId"`
	ChainName  string         `db:"chain_name" json:"chainName"`
	IconUrl    sql.NullString `db:"icon_url" json:"iconUrl"`
	Website    sql.NullString `db:"website" json:"website"`
	InsertedAt time.Time      `db:"inserted_at" json:"insertedAt"`
	UpdatedAt  sql.NullTime   `db:"updated_at" json:"updated_at"`
}

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

// Transaction represents a transaction in cosmos based blockchain
type Transaction struct {
	ID         int64        `db:"id" json:"id"`
	ChainId    int64        `db:"chain_id" json:"chainId"`
	Hash       string       `db:"transaction_hash" json:"hash"`
	Height     BlockHeight  `db:"height" json:"height"`
	Code       int          `db:"code" json:"code"`
	CodeSpace  string       `db:"code_space" json:"codeSpace"`
	TxData     string       `db:"tx_data" json:"txData"`
	RawLog     string       `db:"raw_log" json:"rawLog"`
	Info       string       `db:"info" json:"info"`
	Memo       string       `db:"memo" json:"memo"`
	Seq        int          `db:"seq" json:"seq"`
	GasWanted  uint64       `db:"gas_wanted" json:"gasWanted"`
	GasUsed    uint64       `db:"gas_used" json:"gasUsed"`
	IssuedAt   time.Time    `db:"issued_at" json:"issuedAt"`
	InsertedAt time.Time    `db:"inserted_at" json:"insertedAt"`
	UpdatedAt  sql.NullTime `db:"updated_at" json:"updatedAt"`
}

// Event is emitted when the transaction is executed, it has a one-to-many relationship with the transaction
type Event struct {
	ID         int64        `db:"id" json:"id"`
	ChainId    int64        `db:"chain_id" json:"chainId"`
	TxId       int          `db:"tx_id" json:"txId"`
	Height     BlockHeight  `db:"block_height" json:"height"`
	Seq        uint32       `db:"event_seq" json:"seq"`
	Type       string       `db:"event_type" json:"type"`
	Key        string       `db:"event_key" json:"key"`
	Value      string       `db:"event_value" json:"value"`
	Indexed    bool         `db:"indexed" json:"indexed"`
	InsertedAt time.Time    `db:"inserted_at" json:"insertedAt"`
	UpdatedAt  sql.NullTime `db:"updated_at" json:"updatedAt"`
}

// Message is a message in the transaction, it has a one-to-many relationship with the transaction
type Message struct {
	ID            int64        `db:"id" json:"id"`
	TransactionId int          `db:"transaction_id" json:"transactionId"`
	Seq           int          `db:"seq" json:"seq"`
	RawData       string       `db:"rawdata" json:"rawData"` // RawData is the raw format of message, it's a json string
	InsertedAt    time.Time    `db:"inserted_at" json:"insertedAt"`
	UpdatedAt     sql.NullTime `db:"updated_at" json:"updatedAt"`
}

// Account describes wallet address in the blockchain
type Account struct {
	ID         int64        `db:"id" json:"id"`
	ChainId    int64        `db:"chain_id" json:"chainId"`
	Address    string       `db:"address" json:"address"`
	InsertedAt time.Time    `db:"inserted_at" json:"insertedAt"`
	UpdatedAt  sql.NullTime `db:"updated_at" json:"updatedAt"`
}

// AccountBalance keeps the balance of the certain account
type AccountBalance struct {
	ID         int64        `db:"id" json:"id"`
	AccountId  int64        `db:"account_id" json:"accountId"`
	Amount     uint64       `db:"amount" json:"amount"`
	CoinName   string       `db:"coin_name" json:"coinName"` // CoinName is the name of coin, such as "uatom", actually, it's the denom of coin
	InsertedAt time.Time    `db:"inserted_at" json:"insertedAt"`
	UpdatedAt  sql.NullTime `db:"updated_at" json:"updatedAt"`
}
