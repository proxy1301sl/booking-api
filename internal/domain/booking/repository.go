package booking

import (
	"booking-api/internal/db"

	"context"
)

type Repository struct {
	db *db.Postgres
}

func NewRepository(db *db.Postgres) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Rent(ctx context.Context, resource_id string)(*Flat, error) {
	var res Flat
	query := `SELECT id, user_id, resource_id, start_time, start_time,end_time, created_at FROM bookings WHERE resource_id = $1`
	err := r.db.Pool.QueryRow(ctx, query, resource_id).Scan(&res.id, res.user_id, &res.resource_id, &res.start_time, &res.end_time, &res.created_at )
	if err != nil {
		return nil, err
	}
	return &res, nil
}
