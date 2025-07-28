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
		"id_user": {
			"required": "Id user wajib diisi",
		},
		"title": {
			"required": "Judul acara wajib diisi",
		},
		"date": {
			"required": "Tanggal wajib diisi",
		},
		"time": {
			"required": "Waktu wajib diisi",
		},
		"location": {
			"required": "Lokasi wajib diisi",
		},
		"description": {
			"required": "Deskripsi wajib diisi",
		},
		"primary_color": {
			"required": "Primary color wajib diisi",
		},
		"secondary_color": {
			"required": "Secondary color wajib diisi",
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
