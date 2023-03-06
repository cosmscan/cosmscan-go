package model

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
	"testing"
)

func newRandomEvent() *Event {
	return &Event{
		ChainId: rand.Uint(),
		TxId:    rand.Uint(),
		Height:  rand.Uint32(),
		Seq:     rand.Uint32(),
		Type:    rand.Str(10),
		Key:     rand.Str(10),
		Value:   rand.Str(10),
		Indexed: false,
	}
}

func TestEventCreate(t *testing.T) {
	db := newMemoryDB(t)
	e := newRandomEvent()
	err := e.Create(db)
	require.NoError(t, err)
}

func TestFindAllByTxId(t *testing.T) {
	db := newMemoryDB(t)
	var events []*Event
	txId := rand.Uint()

	for i := 0; i < 10; i++ {
		e := newRandomEvent()
		e.TxId = txId
		err := e.Create(db)
		require.NoError(t, err)
		events = append(events, e)
	}

	e := newRandomEvent()
	events, err := e.FindAllByTxId(db, txId)
	require.NoError(t, err)
	require.Len(t, events, 10)

	// find by unknown tx id
	unknownId := rand.Uint()
	events, err = e.FindAllByTxId(db, unknownId)
	require.NoError(t, err)
	require.Len(t, events, 0)
}
