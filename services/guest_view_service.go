package services

import (
	"be-undangan-digital/config"
	"be-undangan-digital/models"
	"be-undangan-digital/requests"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func CreateGuestViewService(data requests.GuestViewRequest) (*models.GuestView, error) {
	IdGuestView, err := gonanoid.New(8)
	if err != nil {
		return nil, err
	}

	new_guest_view := models.GuestView{
		IdGuestView:      IdGuestView,
		IdInvitationLink: data.IdInvitationLink,
		IpAddress:        data.IpAddres,
		UserAgent:        data.UserAgent,
	}

	err = config.DB.Create(&new_guest_view).Error
	if err != nil {
		return nil, err
	}

	return &new_guest_view, nil
}
