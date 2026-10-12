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

type UserResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
