package models

type SharedSocial struct {
	IdSharedSocial   string `json:"id_shared_social" gorm:"primaryKey;type:char(8)"`
	IdInvitationLink string `json:"id_link" gorm:"type:char;size:8"`
	NamePlatform     string `json:"name_platform" gorm:"type:varchar;size:50"`

	// InvitationLink InvitationLink `gorm:"foreignKey:IdLink;references:IdInvitationLink;"`
}

func (SharedSocial) TableName() string {
	return "shared_socials"
}
