package services

import (
	"be-undangan-digital/config"
	"be-undangan-digital/models"
	"errors"
	"fmt"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

func CreateInvitationLink(IdInvitation string, link string) (*models.InvitationLink, error) {
	var InvitationLinkAlreadyExist models.InvitationLink

	tx := config.DB.Where("id_invitation = ?", IdInvitation).Take(&InvitationLinkAlreadyExist)
	if tx.Error == nil && tx.RowsAffected > 0 {
		return nil, errors.New("tautan undangan sudah ada")
	}
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("gagal mengecek tautan undangan: %w", tx.Error)
	}

	IdInvitationLink, err := gonanoid.New(8)
	if err != nil {
		return nil, err
	}

	new_invitation_link := models.InvitationLink{
		IdInvitationLink: IdInvitationLink,
		IdInvitation:     IdInvitation,
		Link:             link,
		IsActive:         true,
		CreatedAt:        time.Now(),
	}

	if err := config.DB.Create(new_invitation_link).Error; err != nil {
		return nil, err
	}

	return &new_invitation_link, nil
}

func GetInvitationLink(IdInvitation string) (*models.InvitationLink, error) {
	var InvitationLink models.InvitationLink

	err := config.DB.Where("id_invitation = ?", IdInvitation).First(&InvitationLink).Error
	if err != nil {
		return nil, err
	}

	return &InvitationLink, nil
}
