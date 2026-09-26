package config

import (
	"os"
	"strconv"
	"time"
)

type DBConfig struct {
	DatabaseURL string
	MaxOpenConn int
	MaxIdleConn int
}

type Config struct {
	JWTSecret              string
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
	DB                     *DBConfig
}

func Load() (*Config, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	accessTokenExpiration := os.Getenv("ACCESS_TOKEN_EXPIRATION")
	refreshTokenExpiration := os.Getenv("REFRESH_TOKEN_EXPIRATION")
	accessExp, err := time.ParseDuration(accessTokenExpiration)
	if err != nil {
		return nil, err
	}
	refreshExp, err := time.ParseDuration(refreshTokenExpiration)
	if err != nil {
		return nil, err
	}
	MaxOpenConn, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONN"))
	if err != nil {
		return nil, err
	}
	MaxIdleConn, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONN"))
	if err != nil {
		return nil, err
	}

	dbConfig := &DBConfig{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		MaxOpenConn: MaxOpenConn,
		MaxIdleConn: MaxIdleConn,
	}
	return &Config{
		JWTSecret:              jwtSecret,
		AccessTokenExpiration:  accessExp,
		RefreshTokenExpiration: refreshExp,
		DB:                     dbConfig,
	}, nil

}
