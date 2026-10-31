package booking

import (
	"booking-api/internal/domain/httpx"
	"booking-api/internal/domain/responce"
	"booking-api/internal/domain/validator"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	svc *Service
	v   *validator.Validatorr
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
	id, err := h.svc.CreateBooking(c.Request.Context(), id, req.resource_id, req.start_time, req.end_time)
	if err != nil {
		if errors.Is(err, httpx.ErrConflict) {
			responce.ErrorResponse(c, http.StatusConflict, err.Error(), nil)
			return
		}
		responce.ErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)

	}
	responce.SuccessResponse(c, http.StatusCreated, "", id)

}
