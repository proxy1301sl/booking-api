package user

import "time"

type User struct {
	Email        string
	PasswordHash string
	Role         string
	ID           string
	CreatedAt    time.Time
}

type LoginRequest struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}
