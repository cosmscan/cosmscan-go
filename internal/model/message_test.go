package model

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
	"testing"
)

func newRandomMessage(rawData []byte) *Message {
	return &Message{
		TransactionId: rand.Uint(),
		Seq:           rand.Uint(),
		RawData:       rawData,
	}
}

func TestMessageCreate(t *testing.T) {
	db := newMemoryDB(t)
	m := newRandomMessage([]byte("{\"test\": 1}"))
	err := m.Create(db)
	require.NoError(t, err)
}

func TestMessageFindAllByTxId(t *testing.T) {
	db := newMemoryDB(t)

	m := newRandomMessage([]byte("{\"test\": 1}"))
	err := m.Create(db)
	require.NoError(t, err)

	messages, err := m.FindAllByTxId(db, m.TransactionId)
	require.NoError(t, err)
	require.Len(t, messages, 1)
}
