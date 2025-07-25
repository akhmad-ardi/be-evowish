package services

import (
	"be-undangan-digital/config"
	"be-undangan-digital/lib"
	"be-undangan-digital/models"
	"errors"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func CreateUserService(name string, email string, password string) (*models.User, error) {
	var userAlreadyExist models.User

	err := config.DB.Where("email = ?", email).First(&userAlreadyExist).Error
	if err != nil {
		return nil, errors.New("email sudah digunakan")
	}

	passwordHash, err := lib.HashPassword(password)
	if err != nil {
		return nil, err
	}

	id_user, err := gonanoid.New(8)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		IdUser:   id_user,
		Name:     name,
		Email:    email,
		Password: passwordHash,
	}

	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByField(field string, value interface{}) (*models.User, error) {
	var user models.User

	err := config.DB.Where(fmt.Sprintf("%s = ?", field), value).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
