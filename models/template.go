package models

import (
	"time"

	"gorm.io/datatypes"
)

type Template struct {
	IdTemplate      string         `json:"id_template" gorm:"primaryKey;type:char(8)"`
	Name            string         `json:"name" gorm:"type:varchar(50)"`
	TemplateImage   string         `json:"template_image" gorm:"type:varchar(50)"`
	DataTemplate    datatypes.JSON `json:"data_template" gorm:"type:jsonb"`
	BackgroundImage string         `json:"background_image" gorm:"type:varchar(50)"`
	CreatedAt       time.Time      `json:"created_at" gorm:"autoCreateTime:false"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"autoCreateTime:false"`

	Invitations []Invitation `gorm:"foreignKey:IdTemplate;references:IdTemplate;"`
}

func (Template) TableName() string {
	return "templates"
}
