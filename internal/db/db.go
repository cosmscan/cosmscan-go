package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

// New open database connection with given config
func New(cfg Config) (*DB, error) {
	var db *gorm.DB
	var err error

	switch cfg.Driver {
	case "sqlite":
		db, err = openSqlite()
	case "postgresql":
		db, err = openPostgres(cfg)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", cfg.Driver)
	}

	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

func openSqlite() (*gorm.DB, error) {
	dsn := "file:memory:"
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}

func openPostgres(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}

func (d *DB) Instance() *gorm.DB {
	return d.db
}
