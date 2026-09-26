package auth

import (
	"booking-api/internal/domain/user"
	"encoding/json"
	"net/http"
)

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req user.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = Validate(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
	}
}
