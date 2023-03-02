package model

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
	"gorm.io/gorm"
)

func TestBlockCreate(t *testing.T) {
	db := newMemoryDB(t)

	chainId := rand.Uint()
	block := &Block{
		ChainID:            chainId,
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
	err := block.Create(db)
	require.Error(t, err)

	chain := &Chain{
		Model:     gorm.Model{ID: chainId},
		ChainId:   rand.Str(10),
		ChainName: rand.Str(10),
	}
	require.NoError(t, chain.Create(db))

	err = block.Create(db)
	require.NoError(t, err)
}
