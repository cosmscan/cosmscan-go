package model

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAccountBalanceCreate(t *testing.T) {
	db := newMemoryDB(t)
	ab := &AccountBalance{
		AccountID: 1,
		Amount:    1000,
		CoinName:  "uatom",
	}

	err := ab.Create(db)
	require.NoError(t, err)
}
