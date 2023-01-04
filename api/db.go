package api

import (
	"cosmscan-go/config"
	"cosmscan-go/db"
	"cosmscan-go/db/psqldb"

	"go.uber.org/zap"
)

// MustInitDB initializes the database connection.
func MustInitDB(cfg *config.ServerConfig) db.DB {
	database, err := psqldb.NewPsqlDB(&psqldb.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Database: cfg.DB.Database,
	})
	if err != nil {
		zap.S().Fatalw("failed to initialize database", "error", err)
		return nil
	}
	return database
}
