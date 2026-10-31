package booking

import (
	"booking-api/internal/db"
	"booking-api/internal/domain/httpx"
	"time"

	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

type Repository struct {
	db *db.Postgres
}

func NewRepository(db *db.Postgres) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateBooking(userID, resID int64, startTime, endTime *time.Time) (int64, error) {
	ctx := context.Background()
	query := `INSERT INTO booking (user_id, start_time, end_time) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.Pool.QueryRow(ctx, query, userID, startTime, endTime).Scan(&resID)
	pgerr, ok := err.(*pgconn.PgError)
	if !ok {
		return 0, err
	}
	if pgerr.Code == "23P01" {
		return 0, httpx.ErrConflict
	}

	return resID, nil

}
