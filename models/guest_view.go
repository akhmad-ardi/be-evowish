package models

import (
	"time"
)

type GuestView struct {
	IdGuestView      string    `json:"id_guest_view" gorm:"primaryKey;type:char(8)"`
	IdInvitationLink string    `json:"id_link" gorm:"type:char(8)"`
	IpAddress        string    `json:"ip_address" gorm:"type:varchar;size:45"`
	UserAgent        string    `json:"user_agent"`
	ViewedAt         time.Time `json:"viewed_at" gorm:"autoCreateTime:false"`

	// InvitationLink InvitationLink `gorm:"foreignKey:IdLink;references:IdInvitationLink;"`
}

func (GuestView) TableName() string {
	return "guest_views"
}
