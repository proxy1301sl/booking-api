package booking

import (
	"booking-api/internal/db"
	"booking-api/internal/domain/httpx"
	"errors"
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

func (r *Repository) CreateBooking(ctx context.Context, userID, resID int64, startTime, endTime *time.Time) (int64, error) {
	var id int64
	query := `INSERT INTO core.booking (user_id, resource_id, start_time, end_time) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.Pool.QueryRow(ctx, query, userID, resID, startTime, endTime).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23P01" {
				return 0, httpx.ErrConflict
			}
		}
		return 0, err
	}

	return id, nil

}

func (r *Repository) GetBookingByID(ctx context.Context, userID int64) ([]Response, error) {
	rows, err := r.db.Pool.Query(ctx, `SELECT id, user_id, resource_id, start_time, end_time FROM core.booking WHERE user_id = $1 ORDER BY id`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []Response
	for rows.Next() {
		var v Response
		if err = rows.Scan(&v.ID, &v.User_id, &v.Resource_id, &v.Start_time, &v.End_time); err != nil {
			return nil, err
		}
		data = append(data, v)
	}
	return data, rows.Err()

}
