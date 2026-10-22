package resource

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) GetAvailableApartments(ctx context.Context, p *ListParams) ([]Resource, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, name, description, location, is_active, created_at FROM resource WHERE is_active = true LIMIT $1 OFFSET $2", p.Limit, p.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []Resource
	for rows.Next() {
		var w Resource
		if err := rows.Scan(&w.ID, &w.Name, &w.Description, &w.Location, &w.IsActive, &w.CreatedAt); err != nil {
			return nil, err
		}
		data = append(data, Resource{
			ID:          w.ID,
			Name:        w.Name,
			Description: w.Description,
			Location:    w.Location,
			IsActive:    w.IsActive,
			CreatedAt:   w.CreatedAt,
		})
		if err = rows.Err(); err != nil {
			return data, err
		}
	}
	return data, nil
}
