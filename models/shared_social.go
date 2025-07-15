package models

type SharedSocial struct {
	IdSharedSocial string `json:"id_shared_social" gorm:"primaryKey;size:10"`
	IdLink         string `json:"id_link"`
	NamePlatform   string `json:"name_platform"`

	// InvitationLink InvitationLink `gorm:"foreignKey:IdLink;references:IdInvitationLink;"`
}

func (SharedSocial) TableName() string {
	return "shared_socials"
}
