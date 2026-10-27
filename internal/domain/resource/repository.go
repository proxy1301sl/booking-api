package resource

import (
	db2 "booking-api/internal/db"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *db2.Postgres
}

func New(db *db2.Postgres) *Repository {
	return &Repository{db: db}
}

func (r *Repository) ListAvailableResources(ctx context.Context, p *ListParams) ([]Resource, error) {
	rows, err := r.db.Pool.Query(ctx, "SELECT id, name, description, location, is_active, created_at FROM resource WHERE is_active = true LIMIT $1 OFFSET $2", p.Limit, p.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []Resource
	for rows.Next() {
		var w Resource
		if err = rows.Scan(&w.ID, &w.Name, &w.Description, &w.Location, &w.IsActive, &w.CreatedAt); err != nil {
			return nil, err
		}
		data = append(data, w)
	}
	return data, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*ListResource, error) {
	var res ListResource
	query := `SELECT id, name, description, location, is_active, created_at FROM resource WHERE id = $1`
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(&res.ID, &res.Name, &res.Description, &res.Location, &res.IsActive, &res.CreatedAt)
	if err == nil {
		return &res, nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	return nil, err
}
