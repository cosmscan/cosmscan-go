package model

import (
	"cosmscan-go/internal/db"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func newMemoryDB(t *testing.T) *gorm.DB {
	t.Helper()

	d, err := db.New(db.Config{
		Driver: "sqlite",
	})
	require.NoError(t, err, "error while creating sqlite database")

	err = d.AutoMigrate(ModelsToAutoMigrate())
	require.NoError(t, err, "error while migrating schemas")
	return d.Instance()
}
