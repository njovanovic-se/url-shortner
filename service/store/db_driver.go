package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type DatabaseConfig struct {
	MaxIdleConnections    int
	MaxOpenConnections    int
	ConnMaxLifetime       time.Duration
	ConnMaxIdleLifetime   time.Duration
	DefaultContextTimeout time.Duration
}

func (config *DatabaseConfig) ApplyDefaults() {
	if config.ConnMaxIdleLifetime == 0 {
		config.ConnMaxIdleLifetime = 10 * time.Minute
	}

	if config.ConnMaxLifetime == 0 {
		config.ConnMaxLifetime = 5 * time.Minute
	}

	if config.DefaultContextTimeout == 0 {
		config.DefaultContextTimeout = 5 * time.Second
	}
}

func NewPostgresDB(dsn string, dbConfig DatabaseConfig) (*DB, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	db.SetMaxIdleConns(dbConfig.MaxIdleConnections)
	db.SetMaxOpenConns(dbConfig.MaxOpenConnections)
	db.SetConnMaxIdleTime(dbConfig.ConnMaxIdleLifetime)
	db.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), dbConfig.DefaultContextTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Printf("Successfully connected to the database")
	return &DB{DB: db}, nil
}

func (d *DB) Close() error {
	return d.DB.Close()
}
