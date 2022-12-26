package db

import "context"

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

	// InsertMessage inserts the given message into the database.
	InsertMessage(ctx context.Context, message *Message) (int64, error)
}

type BlockReader interface {
	// Block returns the block with the given height.
	Block(ctx context.Context, height BlockHeight) (*Block, error)

	// BlockByHash returns the block with the given hash.
	BlockByHash(ctx context.Context, hash string) (*Block, error)
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

type DB interface {
	Writer
	BlockReader
	TransactionReader
	EventReader
	MessageReader

	// Close closes the database, freeing up any available resources.
	Close() error
}
