package booking

import (
	"booking-api/internal/domain/httpx"
	"booking-api/internal/domain/responce"
	"booking-api/internal/domain/validator"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type BookingHandler struct {
	svc *Service
	v   *validator.Validatorr
}

func NewBookingHandler(svc *Service, v *validator.Validatorr) *BookingHandler {
	return &BookingHandler{svc: svc, v: v}
}

func (h *BookingHandler) Booking(c *gin.Context) {
	var req Request
	if err := c.ShouldBind(&req); err != nil {
		responce.ErrorResponse(c, http.StatusBadRequest, err.Error(), err)
		return
	}
	if validation := h.v.Validate(req); validation != nil {
		responce.ErrorResponse(c, http.StatusBadRequest, validation.Error(), nil)
		return
	}
	if !req.Start_time.Before(*req.End_time) {
		responce.ErrorResponse(c, http.StatusBadRequest, "start_time must be before end_time", nil)
		return
	}
	userID, exist := c.Get("user_id")
	if !exist {
		responce.ErrorResponse(c, http.StatusUnauthorized, "not_found", nil)
		return
	}
	id, ok := userID.(int64)
	if !ok {
		responce.ErrorResponse(c, http.StatusUnauthorized, "not_found", nil)
		return
	}
	id, err := h.svc.CreateBooking(c.Request.Context(), id, req.Resource_id, req.Start_time, req.End_time)
	if err != nil {
		if errors.Is(err, httpx.ErrConflict) {
			responce.ErrorResponse(c, http.StatusConflict, err.Error(), nil)
			return
		}
		if errors.Is(err, pgx.ErrNoRows) {
			responce.ErrorResponse(c, http.StatusNotFound, "resource not found", nil)
			return
		}
		if errors.Is(err, httpx.ErrResourceInactive) {
			responce.ErrorResponse(c, http.StatusConflict, err.Error(), nil)
			return
		}
		responce.ErrorResponse(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}
	responce.SuccessResponse(c, http.StatusCreated, "", id)

}

func (h *BookingHandler) GetBookings(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		responce.ErrorResponse(c, http.StatusUnauthorized, "not_found", nil)
		return
	}
	id, ok := userID.(int64)
	if !ok {
		responce.ErrorResponse(c, http.StatusInternalServerError, "", nil)
		return
	}
	res, err := h.svc.GetBookingByID(c.Request.Context(), id)
	if err != nil {
		responce.ErrorResponse(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}
	responce.SuccessResponse(c, http.StatusOK, "", res)

}
