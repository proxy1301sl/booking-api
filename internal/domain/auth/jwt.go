package auth

import (
	"booking-api/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	config *config.Config
}

type Claims struct {
	UserID    string `json:"user_id"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

func NewManager(config *config.Config) *Manager {
	return &Manager{config: config}
}

//refresh token

func (m *Manager) CreateToken(userID string) (string, error) {
	CustomClaims := Claims{
		UserID:           userID,
		Role:             "",
		TokenType:        "",
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims)
	tokenString, err := token.SignedString([]byte(m.config.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
