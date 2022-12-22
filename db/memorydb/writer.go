package memorydb

import (
	"cosmscan-go/db"

	sq "github.com/Masterminds/squirrel"
)

func (m *MemoryDB) InsertChain(chain *db.Chain) (int64, error) {
	ret, err := sq.Insert("chains").Columns("chain_id", "chain_name", "inserted_at", "updated_at").
		Values(chain.ChainId, chain.ChainName, chain.InsertedAt, chain.UpdatedAt).
		RunWith(m.db).Exec()

	if err != nil {
		return 0, err
	}

	lastId, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (m *MemoryDB) InsertBlock(block *db.Block) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryDB) InsertTransaction(tx *db.Transaction) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryDB) InsertEvent(event *db.Event) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryDB) InsertAccount(account *db.Account) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryDB) InsertMessage(message *db.Message) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemoryDB) InsertOrUpdateAccountBalance(accountBalance *db.AccountBalance) (int64, error) {
	//TODO implement me
	panic("implement me")
}
