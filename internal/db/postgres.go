package db

import (
	"booking-api/internal/config"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	DB *pgxpool.Pool
}

func Connect(cfg *config.DBConfig, ctx context.Context) (*Postgres, error) {
	conf, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	conf.MaxConns = int32(cfg.MaxOpenConn)
	conf.MinConns = int32(cfg.MaxIdleConn)
	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}
	return &Postgres{DB: pool}, nil
}
