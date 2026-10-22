package auth

import (
	"booking-api/internal/config"
	"booking-api/internal/domain/httpx"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	cfg *config.Config
}

type JWTClaims struct {
	UserID int64
	jwt.RegisteredClaims
}

func NewManager(config *config.Config) *Manager {
	return &Manager{cfg: config}
}

func (m *Manager) CreateToken(UserID int64) (string, error) {
	claims := JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(UserID, 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.cfg.AccessTokenExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(m.cfg.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (m *Manager) ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(m.cfg.JWTSecret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, httpx.ErrExpiredToken
		}
		return nil, httpx.ErrInvalidToken
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid

}
