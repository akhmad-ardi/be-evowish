package models

import (
	"time"
)

type Template struct {
	IdTemplate string    `json:"id_template" gorm:"primaryKey;type:char(8)"`
	Name       string    `json:"name" gorm:"type:varchar(50)"`
	PreviewURL string    `json:"preview_url" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime:false"`

	Invitations []Invitation `gorm:"foreignKey:IdTemplate;references:IdTemplate;"`
}

func (Template) TableName() string {
	return "templates"
}
