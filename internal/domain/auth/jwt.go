package auth

import (
	"booking-api/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	config *config.Config
}

type CustomClaims struct {
	UserID    string `json:"user_id"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

func NewManager(config *config.Config) *Manager {
	return &Manager{config: config}
}

func (m *Manager) CreateToken(userID string) (string, error) {
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.config.AccessTokenExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(m.config.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
