package auth

type RegisterRequest struct {
	Email    string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
