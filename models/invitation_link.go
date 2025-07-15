package models

import (
	"time"
)

type InvitationLink struct {
	IdInvitationLink string    `json:"id_invitation_link" gorm:"primaryKey;size:10"`
	IdInvitation     string    `json:"id_invitation"`
	Link             string    `json:"link"`
	IsActive         string    `json:"is_active"`
	CreatedAt        time.Time `json:"created_at"`

	SharedSocial []SharedSocial `gorm:"foreignKey:IdLink;references:IdInvitationLink;"`
	GuestView    []GuestView    `gorm:"foreignKey:IdLink;references:IdInvitationLink;"`
}

func (InvitationLink) TableName() string {
	return "invitation_links"
}
