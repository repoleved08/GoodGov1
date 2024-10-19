package services

import (
	"auth-api/models"
	"auth-api/repositories"
	"auth-api/utils"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func (au *AuthService) CreateUser(username, email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	user := &models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}
return au.UserRepo.CreateUser(user)
}
