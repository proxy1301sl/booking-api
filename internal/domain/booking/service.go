package booking

import (
	"booking-api/internal/domain/httpx"
	"booking-api/internal/domain/resource"
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

type Service struct {
	repo Repository
	res  resource.Repository
}

func NewBookingService(repo Repository, res resource.Repository) *Service {
	return &Service{repo: repo, res: res}
}

func (h *Service) CreateBooking(ctx context.Context, userID, resID int64, startTime *time.Time, endTime *time.Time) (int64, error) {
	res, err := h.res.GetByID(ctx, resID)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, pgx.ErrNoRows
	}
	if err != nil {
		return 0, err
	}
	id, err := h.repo.CreateBooking(userID, res.ID, startTime, endTime)
	if errors.Is(err, httpx.ErrConflict) {
		return 0, err
	}
	return id, nil
}
