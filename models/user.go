package models

import (
	"time"
)

type User struct {
	IdUser    string    `json:"id_user" gorm:"primaryKey;type:char(8)"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime:false"`

	Invitations []Invitation `gorm:"foreignKey:IdUser;references:IdUser;"`
}

func (User) TableName() string {
	return "users"
}
