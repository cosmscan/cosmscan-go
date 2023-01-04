package db

import (
	"context"
)

type Writer interface {
	// InsertChain inserts a chain information into the database.
	InsertChain(ctx context.Context, chain *Chain) (int64, error)

	// InsertBlock inserts the given block into the database.
	InsertBlock(ctx context.Context, block *Block) (int64, error)

	// InsertTransaction inserts the given transaction into the database.
	InsertTransaction(ctx context.Context, tx *Transaction) (int64, error)

	// InsertEvent inserts the given event into the database.
	InsertEvent(ctx context.Context, event *Event) (int64, error)

	// InsertAccount inserts the given account into the database.
	InsertAccount(ctx context.Context, account *Account) (int64, error)

	// UpdateAccountBalance updates the balance of the given account.
	// If the account does not exist, new record is created.
	UpdateAccountBalance(ctx context.Context, accountId int64, coinName string, balance uint64) error

	// InsertMessage inserts the given message into the database.
	InsertMessage(ctx context.Context, message *Message) (int64, error)
}

type BlockReader interface {
	// Block returns the block with the given height.
	Block(ctx context.Context, height BlockHeight) (*Block, error)

	// BlockByHash returns the block with the given hash.
	BlockByHash(ctx context.Context, hash string) (*Block, error)

	// LatestBlock returns the latest saved block in the database
	LatestBlock(ctx context.Context) (*Block, error)
}

type ChainReader interface {
	// AllChains returns all chains in the database.
	AllChains(ctx context.Context) ([]*Chain, error)

	// FindChainByName returns the chain with the given name.
	FindChainByName(ctx context.Context, name string) (*Chain, error)
}

type TransactionReader interface {
	// Transaction returns the transaction with the given hash.
	Transaction(ctx context.Context, hash string) (*Transaction, error)
}

type EventReader interface {
	// EventsInTx returns the events in the given transaction.
	EventsInTx(ctx context.Context, txId int) ([]*Event, error)
}

type MessageReader interface {
	// MessagesInTx returns the messages in the given transaction.
	MessagesInTx(ctx context.Context, txId int) ([]*Message, error)
}

type AccountReader interface {
	// FindAccountByAddress returns whether the account exists or not
	FindAccountByAddress(ctx context.Context, chainId int64, address string) (*Account, error)
}

type DB interface {
	Writer
	BlockReader
	TransactionReader
	ChainReader
	EventReader
	MessageReader
	AccountReader

	// WithTransaction executes the given function within a transaction.
	// if fn returns an error, the transaction is rolled back.
	WithTransaction(ctx context.Context, fn func(tx DB) error) error

	// Close closes the database, freeing up any available resources.
	Close() error
}
