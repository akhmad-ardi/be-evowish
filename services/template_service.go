package services

import (
	"be-undangan-digital/config"
	"be-undangan-digital/models"
	"errors"
)

func GetTemplatesService() (*[]models.Template, error) {
	var templates []models.Template

	err := config.DB.Find(&templates).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &templates, nil
}
