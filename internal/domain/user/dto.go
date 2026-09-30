package user

import "time"

type User struct {
	Email        string
	PasswordHash string
	Role         string
	ID           int64
	CreatedAt    time.Time
	Token        string
}
