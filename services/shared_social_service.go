package services

import (
	"be-undangan-digital/config"
	"be-undangan-digital/models"
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func CreateSharedSocialService(IdInvitationLink string, NamePlatform string) (*models.SharedSocial, error) {
	var SharedSocialAlreadyExist models.SharedSocial

	err := config.DB.Where("id_invitation_link = ? AND name_platform = ?", IdInvitationLink, NamePlatform).First(&SharedSocialAlreadyExist).Error
	if err == nil {
		return nil, errors.New("sudah dibagikan ke platform " + SharedSocialAlreadyExist.NamePlatform)
	}

	IdSharedSocial, err := gonanoid.New(8)
	if err != nil {
		return nil, errors.New("terjadi kesalahan pada generate id")
	}

	new_shared_social := models.SharedSocial{
		IdSharedSocial:   IdSharedSocial,
		IdInvitationLink: IdInvitationLink,
		NamePlatform:     NamePlatform,
	}

	if err := config.DB.Create(new_shared_social).Error; err != nil {
		return nil, err
	}

	return &new_shared_social, nil
}

func GetSharedSocialService(IdInvitation string, NamePlatform string) (*models.SharedSocial, error) {
	var SharedSocial models.SharedSocial

	err := config.DB.Where("id_invitation = ? AND name_platform = ?", IdInvitation, NamePlatform).First(&SharedSocial).Error
	if err != nil {
		return nil, err
	}

	return &SharedSocial, nil
}
