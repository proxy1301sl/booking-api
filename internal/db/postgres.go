package db

import (
	"booking-api/internal/config"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewDB(cfg *config.DBConfig) (*Postgres, error) {
	ctx := context.Background()
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
		pool.Close()
		return nil, err
	}
	return &Postgres{Pool: pool}, nil
}

func (p *Postgres) Close() {
	p.Pool.Close()

}
