package models

import (
	"time"

	"gorm.io/datatypes"
)

type Invitation struct {
	IdInvitation   string         `json:"id_invitation" gorm:"primaryKey;type:char(8)"`
	IdUser         string         `json:"id_user" gorm:"type:char(8)"`
	IdTemplate     string         `json:"id_template" gorm:"type:char(8)"`
	Name           string         `json:"name" gorm:"type:varchar(50)"`
	DataInvitation datatypes.JSON `json:"data_invitation" gorm:"type:jsonb"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime:false"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoCreateTime:false"`

	InvitationLink InvitationLink `gorm:"foreignKey:IdInvitation;references:IdInvitation;constraint:OnDelete:CASCADE"`
}

func (Invitation) TableName() string {
	return "invitations"
}
