package models

import "gorm.io/gorm"

// swagger:model user
type User struct {
	gorm.Model

	// the name for this user
	// required: true
	Name string `gorm:"not null"`

	// required: true
	Email string `gorm:"uniqueIndex;not null"`

	// required: true
	Password string `gorm:"not null"`
}
