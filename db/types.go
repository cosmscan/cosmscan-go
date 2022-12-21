package db

type Writer interface {
	// InsertChain inserts a chain information into the database.
	InsertChain(chainID string) error

	// InsertBlock inserts the given block into the database.
	InsertBlock(block *Block) error

	// InsertTransaction inserts the given transaction into the database.
	InsertTransaction(tx *Transaction) error

	// InsertEvent inserts the given event into the database.
	InsertEvent(event *Event) error

	// InsertAccount inserts the given account into the database.
	InsertAccount(account *Account) error

	// InsertMessage inserts the given message into the database.
	InsertMessage(message *Message) error

	// InsertOrUpdateAccountBalance inserts or updates the given account balance into the database.
	InsertOrUpdateAccountBalance(accountBalance *AccountBalance) error
}

type BlockReader interface {
	// Block returns the block with the given height.
	Block(height BlockHeight) (*Block, error)

	// BlockByHash returns the block with the given hash.
	BlockByHash(hash string) (*Block, error)
}

type TransactionReader interface {
	// Transaction returns the transaction with the given hash.
	Transaction(hash string) (*Transaction, error)
}

type EventReader interface {
	// EventsInTx returns the events in the given transaction.
	EventsInTx(txId int) ([]*Event, error)
}

type MessageReader interface {
	// MessagesInTx returns the messages in the given transaction.
	MessagesInTx(txId int) ([]*Message, error)
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
