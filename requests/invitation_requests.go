package requests

import "mime/multipart"

type CreateInvitationRequest struct {
	IdTemplate      string                 `json:"id_template" validate:"required"`
	Name            string                 `json:"name" validate:"required"`
	DataInvitation  map[string]interface{} `json:"data_invitation" validate:"required"`
	BackgroundImage string                 `json:"background_image" validate:"required"`
}

type AddBackgroundImageRequest struct {
	File *multipart.FileHeader `form:"file"`
}

type AddDataInvitation struct {
	DataInvitation map[string]interface{} `json:"data_invitation" validate:"required"`
}

type GenerateLinkRequest struct {
	IdInvitation string `json:"id_invitation" validate:"required"`
}

type ShareSocialMediaRequest struct {
	IdInvitation string `json:"id_invitation" validate:"required"`
	NamePlatform string `json:"name_platform" validate:"required"`
}

type GuestViewRequest struct {
	IdInvitationLink string `json:"id_invitation_link" validate:"required"`
	IpAddres         string `json:"ip_address" validate:"required"`
	UserAgent        string `json:"user_agent" validate:"required"`
}
