package models

import (
	"time"
)

type InvitationLink struct {
	IdInvitationLink string    `json:"id_invitation_link" gorm:"primaryKey;type:char(8)"`
	IdInvitation     string    `json:"id_invitation" gorm:"type:char(8)"`
	Link             string    `json:"link" gorm:"type:varchar;size:255;"`
	IsActive         bool      `json:"is_active"`
	CreatedAt        time.Time `json:"created_at"`

	SharedSocial []SharedSocial `gorm:"foreignKey:IdInvitationLink;references:IdInvitationLink;constraint:OnDelete:CASCADE"`
	GuestView    []GuestView    `gorm:"foreignKey:IdInvitationLink;references:IdInvitationLink;constraint:OnDelete:CASCADE"`
}

func (InvitationLink) TableName() string {
	return "invitation_links"
}
