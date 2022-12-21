package memorydb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// MemoryDB is in-memory database for this project
// It's used for testing purpose only
type MemoryDB struct {
	db *sql.DB
}

func NewMemoryDB() (*MemoryDB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	return &MemoryDB{db: db}, nil
}

func (m *MemoryDB) Close() error {
	return m.db.Close()
}
