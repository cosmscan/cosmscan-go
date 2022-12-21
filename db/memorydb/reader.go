package memorydb

import "cosmscan-go/db"

func (m *MemoryDB) Block(height db.BlockHeight) (*db.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryDB) BlockByHash(hash string) (*db.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryDB) Transaction(hash string) (*db.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryDB) EventsInTx(txId int) ([]*db.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryDB) MessagesInTx(txId int) ([]*db.Message, error) {
	//TODO implement me
	panic("implement me")
}
