package model

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
	"gorm.io/gorm"
)

func newTestBlock() *Block {
	return &Block{
		ChainID:            rand.Uint(),
		Height:             BlockHeight(rand.Intn(10)),
		Hash:               rand.Str(10),
		ParentHash:         rand.Str(10),
		ProposerAddress:    rand.Str(10),
		LastCommitHash:     rand.Str(10),
		DataHash:           rand.Str(10),
		ValidatorsHash:     rand.Str(10),
		NextValidatorsHash: rand.Str(10),
		ConsensusHash:      rand.Str(10),
		AppHash:            rand.Str(10),
		LastResultHash:     rand.Str(10),
		EvidenceHash:       rand.Str(10),
	}
}

func TestBlockCreate(t *testing.T) {
	db := newMemoryDB(t)

	block := newTestBlock()
	err := block.Create(db)
	require.Error(t, err)

	chain := &Chain{
		Model:     gorm.Model{ID: block.ChainID},
		ChainId:   rand.Str(10),
		ChainName: rand.Str(10),
	}
	require.NoError(t, chain.Create(db))

	err = block.Create(db)
	require.NoError(t, err)
}

func TestBlockFindByHash(t *testing.T) {
	db := newMemoryDB(t)
	block := newTestBlock()
	chain := &Chain{
		Model:     gorm.Model{ID: block.ChainID},
		ChainId:   rand.Str(10),
		ChainName: rand.Str(10),
	}
	require.NoError(t, chain.Create(db))

	err := block.Create(db)
	require.NoError(t, err)

	// try with unknown block
	b := &Block{}
	err = b.FindByHash(db, "unknown")
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)

	// find created block
	err = b.FindByHash(db, block.Hash)

	block.Model = b.Model // set model to compare
	require.NoError(t, err)
	require.Equal(t, block, b)
}
