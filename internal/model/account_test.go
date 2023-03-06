package model

import (
	"github.com/stretchr/testify/require"
	trand "github.com/tendermint/tendermint/libs/rand"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func newRandomAccount() *Account {
	return &Account{
		ChainId: uint(rand.Uint32()),
		Address: trand.Str(20),
	}
}

func TestAccountCreate(t *testing.T) {
	db := newMemoryDB(t)
	acc := newRandomAccount()
	err := acc.Create(db)
	require.NoError(t, err)
}

func TestAccountFindBy(t *testing.T) {
	db := newMemoryDB(t)
	acc := newRandomAccount()
	err := acc.Create(db)
	require.NoError(t, err)

	acc2 := newRandomAccount()
	err = acc2.FindBy(db, acc.Address, acc.ChainId)
	require.NoError(t, err)
	require.Equal(t, acc.ID, acc2.ID)

	acc3 := newRandomAccount()
	err = acc3.FindBy(db, trand.Str(10), uint(rand.Uint32()))
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func TestAccountFindByID(t *testing.T) {
	db := newMemoryDB(t)
	acc := newRandomAccount()
	err := acc.Create(db)
	require.NoError(t, err)

	// find by id
	acc2 := newRandomAccount()
	err = acc2.FindByID(db, acc.ID)
	require.NoError(t, err)
	require.Equal(t, acc.ChainId, acc2.ChainId)
	require.Equal(t, acc.Address, acc2.Address)

	// find by unknown id
	acc3 := newRandomAccount()
	err = acc3.FindByID(db, uint(rand.Uint32()))
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}
