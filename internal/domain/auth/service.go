package auth

import (
	"errors"
	"net/mail"
	"unicode/utf8"
)

func Validate(req *LoginRequest) error {
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
