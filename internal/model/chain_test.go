package model

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
	"testing"
)

func newRandomChain() *Chain {
	return &Chain{
		ChainId:   rand.Str(10),
		ChainName: rand.Str(10),
	}
}

func TestChainCreate(t *testing.T) {
	db := newMemoryDB(t)
	c := newRandomChain()
	err := c.Create(db)
	require.NoError(t, err)
}

func TestChainFindByID(t *testing.T) {
	db := newMemoryDB(t)
	c := newRandomChain()
	err := c.Create(db)
	require.NoError(t, err)

	// find by id
	c2 := newRandomChain()
	err = c2.FindByID(db, c.ID)
	require.NoError(t, err)
	require.Equal(t, c.ID, c2.ID)
	require.Equal(t, c.ChainId, c2.ChainId)
	require.Equal(t, c.ChainName, c2.ChainName)

	// find by unknown id
	c3 := newRandomChain()
	err = c3.FindByID(db, rand.Uint())
	require.Error(t, err)
}

func TestChainFindByChainID(t *testing.T) {
	db := newMemoryDB(t)
	c := newRandomChain()
	err := c.Create(db)
	require.NoError(t, err)

	// find by chainId
	c2 := newRandomChain()
	err = c2.FindByChainID(db, c.ChainId)
	require.NoError(t, err)
	require.Equal(t, c.ID, c2.ID)
	require.Equal(t, c.ChainId, c2.ChainId)
	require.Equal(t, c.ChainName, c2.ChainName)

	// find by unknown chainId
	c3 := newRandomChain()
	err = c3.FindByChainID(db, rand.Str(10))
	require.Error(t, err)
}

func TestChainFindAll(t *testing.T) {
	db := newMemoryDB(t)
	var chains []*Chain
	for i := 0; i < 10; i++ {
		c := newRandomChain()
		err := c.Create(db)
		require.NoError(t, err)
		chains = append(chains, c)
	}

	c := newRandomChain()
	allChains, err := c.FindAll(db)
	require.NoError(t, err)
	require.Equal(t, len(chains), len(allChains))
	for i := 0; i < len(chains); i++ {
		require.Equal(t, chains[i].ID, allChains[i].ID)
		require.Equal(t, chains[i].ChainId, allChains[i].ChainId)
		require.Equal(t, chains[i].ChainName, allChains[i].ChainName)
	}
}
