package auth

import (
	"booking-api/internal/domain/responce"
	"booking-api/internal/domain/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
	v   *validator.Validatorr
}

func NewHandler(svc *Service, v *validator.Validatorr) *Handler {
	return &Handler{svc: svc, v: v}
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responce.ErrorResponse(c, http.StatusBadRequest, err.Error(), err)
		return
	}
	if validation := h.v.Validate(req); validation != nil {
		responce.ErrorResponse(c, http.StatusBadRequest, validation.Error(), nil)
		return
	}
	user, err := h.svc.CreateUser(c.Request.Context(), req)
	if err != nil {
		responce.ErrorResponse(c, http.StatusConflict, err.Error(), nil)
		return
	}
	responce.SuccessResponse(c, http.StatusCreated, "User created successfully", user)
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responce.ErrorResponse(c, http.StatusBadRequest, err.Error(), err)
		return
	}
	if validationErr := h.v.Validate(req); validationErr != nil {
		responce.ErrorResponse(c, http.StatusBadRequest, validationErr.Error(), nil)
		return
	}
	token, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		responce.ErrorResponse(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}
	responce.SuccessResponse(c, http.StatusOK, "Token created successfully", token)
}

func (h *Handler) Me(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		responce.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	id, ok := userID.(int64)
	if !ok {
		responce.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	usr, err := h.svc.GetUserID(c, id)
	if err != nil {
		responce.ErrorResponse(c, http.StatusInternalServerError, "", err.Error())
		return
	}
	responce.SuccessResponse(c, http.StatusOK, "User fetched successfully", usr)
}
