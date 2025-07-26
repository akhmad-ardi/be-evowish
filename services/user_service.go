package services

import (
	"be-undangan-digital/config"
	"be-undangan-digital/lib"
	"be-undangan-digital/models"
	"be-undangan-digital/requests"
	"errors"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

func CreateUserService(data *requests.RegisterRequest) (*models.User, error) {
	var userAlreadyExist models.User

	tx := config.DB.Where("email = ?", data.Email).Take(&userAlreadyExist)
	if tx.Error == nil && tx.RowsAffected > 0 {
		return nil, errors.New("email sudah digunakan")
	}
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("gagal mengecek email: %w", tx.Error)
	}

	passwordHashed, err := lib.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	id_user, err := gonanoid.New(8)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		IdUser:   id_user,
		Name:     data.Name,
		Email:    data.Email,
		Password: passwordHashed,
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
