package auth

import (
	"booking-api/internal/domain/user"
	"context"
	"errors"
	"net/mail"
	"time"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrUsernameExists     = errors.New("username already exists")
)

type Service struct {
	repo user.Repository
}

func Validate(req LoginRequest) error {
	if req.Email == "" || req.Password == "" {
		return errors.New("email or password is empty")
	}
	if utf8.RuneCountInString(req.Password) < 8 {
		return errors.New("password is too short")
	}
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return errors.New("invalid email")
	}
	return nil
}

func (s *Service) Register(ctx context.Context, req LoginRequest) (*user.User, error) {
	existing, _ := s.repo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, ErrEmailAlreadyExists
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	usr := user.User{
		Email:        req.Email,
		PasswordHash: string(passwordHash),
		Role:         "user",
		ID:           0,
		CreatedAt:    time.Time{},
	}
	err = s.repo.Create(ctx, &usr)
	if err != nil {
		return nil, err
	}
	return &usr, nil

}
