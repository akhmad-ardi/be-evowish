package models

import (
	"time"
)

type GuestView struct {
	IdGuestView string    `json:"id_guest_view" gorm:"primaryKey;size:10"`
	IdLink      string    `json:"id_link"`
	IpAddress   string    `json:"ip_address"`
	UserAgent   string    `json:"user_agent"`
	ViewedAt    time.Time `json:"viewed_at" gorm:"autoCreateTime:false"`

	// InvitationLink InvitationLink `gorm:"foreignKey:IdLink;references:IdInvitationLink;"`
}

func (GuestView) TableName() string {
	return "guest_views"
}
