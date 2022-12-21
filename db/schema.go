package db

import (
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
	ID         int64     `db:"id"`
	ChainId    string    `db:"chain_id"`
	ChainName  string    `db:"chain_name"`
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// Block represents a block in cosmos based blockchain
type Block struct {
	ID                 int64       `db:"id"`
	ChainId            int         `db:"chain_id"`
	Height             BlockHeight `db:"height"`
	Hash               string      `db:"block_hash"`
	ParentHash         string      `db:"prev_hash"`
	ProposerAddress    string      `db:"proposer_address"`
	LastCommitHash     string      `db:"last_commit_hash"`
	DataHash           string      `db:"data_hash"`
	ValidatorsHash     string      `db:"validators_hash"`
	NextValidatorsHash string      `db:"next_validators_hash"`
	ConsensusHash      string      `db:"consensus_hash"`
	AppHash            string      `db:"app_hash"`
	LastResultHash     string      `db:"last_result_hash"`
	EvidenceHash       string      `db:"evidence_hash"`
	BlockTime          time.Time   `db:"block_time"`
	InsertedAt         time.Time   `db:"inserted_at"`
	UpdatedAt          time.Time   `db:"updated_at"`
}

// Transaction represents a transaction in cosmos based blockchain
type Transaction struct {
	ID         int64       `db:"id"`
	ChainId    int         `db:"chain_id"`
	Hash       string      `db:"transaction_hash"`
	Height     BlockHeight `db:"height"`
	Code       int         `db:"code"`
	CodeSpace  string      `db:"code_space"`
	TxData     string      `db:"tx_data"`
	RawLog     string      `db:"raw_log"`
	Info       string      `db:"info"`
	Memo       string      `db:"memo"`
	GasWanted  uint64      `db:"gas_wanted"`
	GasUsed    uint64      `db:"gas_used"`
	IssuedAt   time.Time   `db:"issued_at"`
	InsertedAt time.Time   `db:"inserted_at"`
	UpdatedAt  time.Time   `db:"updated_at"`
}

// Event is emitted when the transaction is executed, it has a one-to-many relationship with the transaction
type Event struct {
	ID         int64       `db:"id"`
	ChainId    int         `db:"chain_id"`
	TxId       int         `db:"tx_id"`
	TxType     int         `db:"tx_type"` // TxType means the type of transaction, there are three types of tx: 1. NormalTx, 2. BeginBlock, 3. EndBlock
	Height     BlockHeight `db:"block_height"`
	Seq        int         `db:"event_seq"`
	Type       string      `db:"event_type"`
	Key        string      `db:"event_key"`
	Value      string      `db:"event_value"`
	Indexed    bool        `db:"indexed"`
	InsertedAt time.Time   `db:"inserted_at"`
	UpdatedAt  time.Time   `db:"updated_at"`
}

// Message is a message in the transaction, it has a one-to-many relationship with the transaction
type Message struct {
	ID            int64     `db:"id"`
	TransactionId int       `db:"transaction_id"`
	Seq           int       `db:"message_seq"`
	RawData       string    `db:"raw_data"` // RawData is the raw format of message, it's a json string
	InsertedAt    time.Time `db:"inserted_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

// Account describes wallet address in the blockchain
type Account struct {
	ID         int64     `db:"id"`
	ChainId    int       `db:"chain_id"`
	Address    string    `db:"address"`
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// AccountBalance keeps the balance of the certain account
type AccountBalance struct {
	ID         int64     `db:"id"`
	AccountId  int       `db:"account_id"`
	Amount     uint64    `db:"amount"`
	CoinName   string    `db:"coin_name"` // CoinName is the name of coin, such as "uatom", actually, it's the denom of coin
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
