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
	query := `INSERT INTO core.users (email, password_hash) VALUES ($1, $2) RETURNING id, role`
	err := r.db.Pool.QueryRow(ctx, query, u.Email, u.PasswordHash).Scan(&u.ID, &u.Role)
	return err
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	query := `SELECT id, password_hash, email, role, created_at FROM core.users WHERE email = $1`
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
	query := `SELECT id, password_hash, email, role, created_at FROM core.users WHERE id = $1`
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(&u.ID, &u.PasswordHash, &u.Email, &u.Role, &u.CreatedAt)
	if err == nil {
		return &u, nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, httpx.ErrUserNotFound
	}
	return nil, err
}
