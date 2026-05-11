package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // file source
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/adocoder12/Costabackend/internal/config"
)

// NewPool creates a pgxpool with explicit connection limits from config.
// The caller is responsible for calling pool.Close() on shutdown.
func NewPool(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("db - parse config: %w", err)
	}

	// Explicit pool limits — prevents exhausting DB connections under load
	poolConfig.MaxConns = cfg.MaxConns
	poolConfig.MinConns = cfg.MinConns
	poolConfig.MaxConnLifetime = 30 * time.Minute
	poolConfig.MaxConnIdleTime = 5 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("db - new pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("db - ping: %w", err)
	}

	return pool, nil
}

// Migrate runs all pending up migrations from internal/db/migrations.
// Safe to call on every startup — already-applied migrations are skipped.
func Migrate(dsn string) error {
	m, err := migrate.New(
		"file://internal/db/migrations",
		dsn,
	)
	if err != nil {
		return fmt.Errorf("db - migrate new: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("db - migrate up: %w", err)
	}

	return nil
}
