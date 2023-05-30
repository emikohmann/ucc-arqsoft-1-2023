package services

import (
	"errors"
	"go-api/domain"
)

func Auth(user domain.User) (domain.AuthResponse, error) {
	token, err := GenerateToken(user.Username)
	if err != nil {
		return domain.AuthResponse{}, errors.New("forbidden")
	}

	return domain.AuthResponse{
		Token: token,
	}, nil
}

func GenerateToken(username string) (string, error) {

	if username == "emi" {
		return "", errors.New("invalido")
	}

	// Completar
	return "abc123456", nil
}
