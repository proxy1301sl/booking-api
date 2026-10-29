package booking

import (
	"booking-api/internal/domain/responce"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
}

func (h *BookingHandler) Booking(c *gin.Context) {
	var req Request
	if err := c.ShouldBind(&req); err != nil {
		responce.ErrorResponse(c, http.StatusBadRequest, err.Error(), err)
		return
	}
	for i := 0; range


}
