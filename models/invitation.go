package models

import (
	"time"
)

type Invitation struct {
	IdInvitation    string    `json:"id_invitation" gorm:"primaryKey;type:char(8)"`
	IdUser          string    `json:"id_user" gorm:"type:char(8)"`
	IdTemplate      string    `json:"id_template" gorm:"type:char(8)"`
	Title           string    `json:"title" gorm:"type:varchar(50)"`
	Date            time.Time `json:"date" gorm:"autoCreateTime:false"`
	Time            time.Time `json:"time" gorm:"type:time"`
	Location        string    `json:"location" gorm:"type:varchar(50)"`
	Description     string    `json:"description" gorm:"type:varchar(255)"`
	PrimaryColor    string    `json:"primary_color" gorm:"type:char(10)"`
	SecondaryColor  string    `json:"secondary_color" gorm:"type:char(10)"`
	BackgroundImage string    `json:"background_image" gorm:"type:varchar(255)"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime:false"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoCreateTime:false"`

	// User           User           `gorm:"foreignKey:IdUser;references:IdUser"`
	// Template       Template       `gorm:"foreignKey:IdTemplate;references:IdTemplate"`
	InvitationLink InvitationLink `gorm:"foreignKey:IdInvitation;references:IdInvitation"`
}

func (Invitation) TableName() string {
	return "invitations"
}
