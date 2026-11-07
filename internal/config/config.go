package config

import (
	"errors"
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
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET is required")
	}
	accessTokenExpiration := os.Getenv("ACCESS_TOKEN_EXPIRATION")
	accessExp, err := time.ParseDuration(accessTokenExpiration)
	if err != nil {
		return nil, err
	}
	if accessExp <= 0 {
		return nil, errors.New("ACCESS_TOKEN_EXPIRATION must be positive")
	}
	MaxOpenConn, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONN"))
	if err != nil {
		return nil, err
	}
	MaxIdleConn, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONN"))
	if err != nil {
		return nil, err
	}
	if MaxIdleConn > MaxOpenConn {
		MaxIdleConn = MaxOpenConn
	}
	dbConfig := &DBConfig{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		MaxOpenConn: MaxOpenConn,
		MaxIdleConn: MaxIdleConn,
	}
	return &Config{
		JWTSecret:             jwtSecret,
		AccessTokenExpiration: accessExp,
		DB:                    dbConfig,
	}, nil

}
