package validations

import (
	"be-undangan-digital/lib"
	"be-undangan-digital/requests"
)

func ValidateCreateInvitationRequest(data requests.CreateInvitationRequest) map[string]string {
	messages := map[string]map[string]string{
		"id_template": {
			"required": "Id template wajib diisi",
		},
		"name": {
			"required": "Nama acara wajib diisi",
		},
		"data_invitation": {
			"required": "Data invitation wajib diisi",
		},
		"background_image": {
			"required": "Background image wajib diisi",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}

func ValidateAddDataInvitation(data requests.AddDataInvitation) map[string]string {
	messages := map[string]map[string]string{
		"data_invitation": {
			"required": "Silahkan isi data invitation",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}

func ValidateAddBackgroundImageRequest(data requests.AddBackgroundImageRequest) map[string]string {
	messages := map[string]map[string]string{
		"file": {
			"required": "File gambar wajib diunggah",
			"image":    "File harus berupa gambar (jpg, jpeg, png, gif)",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}

func ValidateGenerateLinkRequest(data requests.GenerateLinkRequest) map[string]string {
	messages := map[string]map[string]string{
		"id_invitation": {
			"required": "Id invitation wajib diisi",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}

func ValidateShareSocialMediaRequest(data requests.ShareSocialMediaRequest) map[string]string {
	messages := map[string]map[string]string{
		"id_invitation": {
			"required": "Id invitation wajib diisi",
		},
		"name_platform": {
			"required": "Nama platform wajib diisi",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}

func ValidateGuestViewRequest(data requests.GuestViewRequest) map[string]string {
	messages := map[string]map[string]string{
		"id_invitation_link": {
			"required": "Id invitation link wajib diisi",
		},
		"ip_address": {
			"required": "Ip address wajib diisi",
		},
		"user_agent": {
			"required": "User agent wajib diisi",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}
