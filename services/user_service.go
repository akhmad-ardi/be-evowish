package services

import (
	"be-evowish/config"
	"be-evowish/lib"
	"be-evowish/models"
	"errors"
)

func RegisterUserService(name string, email string, password string) (*models.User, error) {
	var existing models.User
	if err := config.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return nil, errors.New("email already used")
	}

	passwordHash, _ := lib.HashPassword(password)

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: passwordHash,
	}

	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
