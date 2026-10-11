package model

import (
	"booking-api/internal/domain/httpx"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) Create(ctx context.Context, u *User) error {
	query := `INSERT INTO users (email, password_hash, role) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := r.pool.QueryRow(ctx, query, u.Email, u.PasswordHash, u.Role).Scan(&u.ID, &u.CreatedAt)
	return err
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	query := `SELECT * FROM users WHERE email = $1`
	err := r.pool.QueryRow(ctx, query, email).Scan(&u.ID, &u.PasswordHash, &u.Email, &u.Role, &u.CreatedAt)
	if err == nil {
		return &u, nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, httpx.ErrNotFound
	}

	return nil, err
}
