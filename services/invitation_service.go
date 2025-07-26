package services

import (
	"be-undangan-digital/config"
	"be-undangan-digital/models"
	"be-undangan-digital/requests"
	"errors"
	"fmt"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

func CreateInvitationService(IdUser string, data *requests.CreateInvitationRequest) (*models.Invitation, error) {
	id_invitation, err := gonanoid.New(8)
	if err != nil {
		return nil, errors.New("id not generate")
	}

	parsedDate, err := time.Parse("2006-01-02", data.Date)
	if err != nil {
		return nil, errors.New("parse date is error")
	}

	parseTime, err := time.Parse("15:04:05", data.Time)
	if err != nil {
		return nil, errors.New("parse date is error")
	}

	new_invitation := &models.Invitation{
		IdInvitation:    id_invitation,
		IdTemplate:      data.IdTemplate,
		IdUser:          IdUser,
		Title:           data.Title,
		Date:            parsedDate,
		Time:            parseTime,
		Location:        data.Location,
		Description:     data.Description,
		PrimaryColor:    data.PrimaryColor,
		SecondaryColor:  data.SecondaryColor,
		BackgroundImage: data.BackgroundImage,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := config.DB.Create(new_invitation).Error; err != nil {
		return nil, err
	}

	return new_invitation, nil
}

func GetInvitations(IdUser string) (*[]models.Invitation, error) {
	var invitations []models.Invitation

	err := config.DB.Where("id_user = ?", IdUser).Find(&invitations).Error
	if err != nil {
		return nil, err
	}

	return &invitations, nil
}

func GetInvitation(IdInvitation string) (*models.Invitation, error) {
	var invitation models.Invitation

	println(IdInvitation)
	err := config.DB.Where("id_invitation = ?", IdInvitation).First(&invitation).Error
	if err != nil {
		println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("data undangan tidak ditemukan")
		}
		return nil, err
	}

	return &invitation, nil
}
