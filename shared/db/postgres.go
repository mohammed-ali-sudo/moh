// internal/db/conn.go
// shared/db/postgres.go
package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Open(ctx context.Context, connStr string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("parse conn string: %w", err)
	}

	// Optional tuning
	cfg.MinConns = 1
	cfg.MaxConns = 10
	cfg.HealthCheckPeriod = 30 * time.Second
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.MaxConnIdleTime = 10 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	// Verify
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	fmt.Println("✅ Database connection verified.")
	return pool, nil
}

func MustOpen(ctx context.Context, connStr string) *pgxpool.Pool {
	pool, err := Open(ctx, connStr)
	if err != nil {
		log.Fatalf("❌ open db: %v", err)
	}
	return pool
}
