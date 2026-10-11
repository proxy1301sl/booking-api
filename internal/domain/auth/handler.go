package auth

import (
	validator2 "booking-api/internal/validator"
	"encoding/json"
	"net/http"
)

type Handler struct {
	svc *Service
	v   *validator2.Validator
}

func NewHandler(svc *Service, v *validator2.Validator) *Handler {
	return &Handler{svc: svc, v: v}
}

func Register(h *Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if errors := h.v.Validate(req); errors != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = h.svc.CreateUser(r.Context(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)

	}
}

func Login(h *Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if errors := h.v.Validate(req); errors != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := h.svc.Login(r.Context(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(token)

	}
}
