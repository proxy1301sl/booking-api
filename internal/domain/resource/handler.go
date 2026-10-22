package resource

import (
	"booking-api/internal/domain/responce"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Resources(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "2"))
	if err != nil || page <= 1 {
		responce.ErrorResponse(c, http.StatusBadRequest, "invalid page param", nil)
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit >= 100 {
		responce.ErrorResponse(c, http.StatusBadRequest, "invalid limit param", nil)
		return
	}
	res, err := h.svc.List(c, &Pagination{Page: page, Limit: limit})
	if err != nil {
		responce.ErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	responce.SuccessResponse(c, http.StatusOK, "available resource", res)

}
