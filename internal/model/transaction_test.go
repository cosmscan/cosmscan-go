package model

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
	"gorm.io/gorm"
	"testing"
	"time"
)

func newRandomTransaction() *Transaction {
	return &Transaction{
		ChainId:   rand.Uint(),
		Hash:      rand.Str(10),
		Height:    rand.Uint32(),
		Code:      0,
		CodeSpace: rand.Str(100),
		TxData:    rand.Str(100),
		RawLog:    []byte("{\"item\": 100}"),
		Info:      rand.Str(100),
		Memo:      rand.Str(100),
		Seq:       0,
		GasWanted: 0,
		GasUsed:   0,
		IssuedAt:  time.Now(),
	}
}

func TestTxCreate(t *testing.T) {
	db := newMemoryDB(t)
	tx := newRandomTransaction()
	err := tx.Create(db)
	require.NoError(t, err)
}

func TestTxFindByHash(t *testing.T) {
	db := newMemoryDB(t)
	tx := newRandomTransaction()
	err := tx.Create(db)
	require.NoError(t, err)

	tx2 := newRandomTransaction()
	err = tx2.FindByHash(db, tx.ChainId, tx.Hash)
	require.NoError(t, err)

	// It's necessary to set the model to the original model
	require.Equal(t, tx2.ChainId, tx.ChainId)
	require.Equal(t, tx2.Hash, tx.Hash)

	// find by unknown hash
	err = tx.FindByHash(db, tx.ChainId, "not-exist-hash")
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func TestTxFindAllByHeight(t *testing.T) {
	db := newMemoryDB(t)
	chainId := rand.Uint()
	height := rand.Uint32()

	var txes []*Transaction
	for i := 0; i < 100; i++ {
		tx := newRandomTransaction()
		tx.ChainId = chainId
		tx.Height = height
		tx.Seq = i
		txes = append(txes, tx)
		err := tx.Create(db)
		require.NoError(t, err)
	}

	got, err := new(Transaction).FindAllByHeight(db, chainId, height)
	require.NoError(t, err)
	require.Len(t, got, 100)
}
