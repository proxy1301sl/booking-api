package user

import (

	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo Repository
}

func (s *UserService) CreateUser(ctx context.Context, req LoginRequest) error {
	existing, _ := s.repo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := User{
		Email:        req.Email,
		PasswordHash: string(passwordHash),
		Role:         "user",
		ID:           "",
		CreatedAt:    time.Time{},
	}
	err = s.repo.Create(ctx, &user)
	if err != nil {
		return err
	}
	return nil

}
