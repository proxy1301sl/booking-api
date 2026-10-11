package model

import "time"

type User struct {
	Email        string
	PasswordHash string
	Role         string
	ID           int64
	CreatedAt    time.Time
	Token        string
}

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required"`
}
