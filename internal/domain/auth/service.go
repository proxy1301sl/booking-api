package auth

import (
	"booking-api/internal/domain/model"
	"context"
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"booking-api/internal/domain/httpx"
)

type Service struct {
	repo    model.Repository
	manager Manager
}

func NewService(repo *model.Repository, manager *Manager) *Service {
	return &Service{repo: *repo, manager: *manager}
}

func (s *Service) CreateUser(ctx context.Context, req RegisterRequest) error {
	_, err := s.repo.FindByEmail(ctx, req.Email)
	if err == nil {
		return httpx.ErrEmailAlreadyExists
	}
	if !errors.Is(err, httpx.ErrNotFound) {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	usr := model.User{
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         "user",
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
		return "", httpx.ErrInvalidCredentials
	}
	err = bcrypt.CompareHashAndPassword([]byte(exst.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", httpx.ErrInvalidCredentials
	}
	token, err := s.manager.CreateToken(strconv.FormatInt(exst.ID, 10))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *Service) GetUserID(ctx context.Context, userID int64) (*model.UserResponse, error) {
	user, err := s.repo.FindByID(ctx, strconv.FormatInt(userID, 10))
	if err != nil {
		return nil, err
	}
	usr := model.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
	return &usr, nil
}
