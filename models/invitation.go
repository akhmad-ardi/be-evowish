package models

import (
	"time"
)

type Invitation struct {
	IdInvitation string    `json:"id_invitation" gorm:"primaryKey;size:10"`
	IdUser       string    `json:"id_user"`
	IdTemplate   string    `json:"id_template"`
	Title        string    `json:"title"`
	Date         time.Time `json:"date" gorm:"autoCreateTime:false"`
	Location     string    `json:"location"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime:false"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoCreateTime:false"`

	// User           User           `gorm:"foreignKey:IdUser;references:IdUser"`
	// Template       Template       `gorm:"foreignKey:IdTemplate;references:IdTemplate"`
	InvitationLink InvitationLink `gorm:"foreignKey:IdInvitation;references:IdInvitation"`
}

func (Invitation) TableName() string {
	return "invitations"
}
