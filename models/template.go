package models

import (
	"time"
)

type Template struct {
	IdTemplate string    `json:"id_template" gorm:"primaryKey;size:10"`
	Name       string    `json:"name"`
	PreviewURL string    `json:"preview_url"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime:false"`

	Invitations []Invitation `gorm:"foreignKey:IdTemplate;references:IdTemplate;"`
}

func (Template) TableName() string {
	return "templates"
}
