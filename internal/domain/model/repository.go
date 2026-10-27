package model

import (
	db2 "booking-api/internal/db"
	"booking-api/internal/domain/httpx"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *db2.Postgres
}

func NewUserRepo(db *db2.Postgres) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, u *User) error {
	query := `INSERT INTO users (email, password_hash, role) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := r.db.Pool.QueryRow(ctx, query, u.Email, u.PasswordHash, u.Role).Scan(&u.ID, &u.CreatedAt)
	return err
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	query := `SELECT id, password_hash, email, role, created_at FROM users WHERE email = $1`
	err := r.db.Pool.QueryRow(ctx, query, email).Scan(&u.ID, &u.PasswordHash, &u.Email, &u.Role, &u.CreatedAt)
	if err == nil {
		return &u, nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, httpx.ErrNotFound
	}

	return nil, err
}

func (r *Repository) FindByID(ctx context.Context, id int64) (*User, error) {
	var u User
	query := `SELECT id, password_hash, email, role, created_at FROM users WHERE id = $1`
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(&u.ID, &u.PasswordHash, &u.Email, &u.Role, &u.CreatedAt)
	if err == nil {
		return &u, nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, httpx.ErrUserNotFound
	}
	return nil, err
}
