package auth

import (
	"booking-api/internal/domain/user"
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrUsernameExists     = errors.New("username already exists")
	ErrNotFound           = errors.New("not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Service struct {
	repo    user.Repository
	manager Manager
}

func NewService(repo *user.Repository, manager *Manager) *Service {
	return &Service{repo: *repo, manager: *manager}
}

func (s *Service) Register(ctx context.Context, req RegisterRequest) error {
	_, err := s.repo.FindByEmail(ctx, req.Email)
	if err == nil {
		return ErrEmailAlreadyExists
	}
	if errors.Is(err, ErrNotFound) {
		return ErrNotFound
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	usr := user.User{
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         "user",
		ID:           0,
		CreatedAt:    time.Time{},
	}
	err = s.repo.Create(ctx, &usr)
	if err != nil {
		return err
	}
	return nil

}

func (s *Service) Login(ctx context.Context, req LoginRequest) (string, error) {
	exst, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", ErrInvalidCredentials
	}
	err = bcrypt.CompareHashAndPassword([]byte(exst.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", ErrInvalidCredentials
	}
	token, err := s.manager.CreateToken(exst)
	if err != nil {
		return "", err
	}
	return token, nil
}
