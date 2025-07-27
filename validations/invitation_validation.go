package validations

import (
	"be-undangan-digital/lib"
	"be-undangan-digital/requests"
)

func ValidateCreateInvitationRequest(data requests.CreateInvitationRequest) map[string]string {
	messages := map[string]map[string]string{
		"IdTemplate": {
			"required": "Id template wajib diisi",
		},
		"IdUser": {
			"required": "Id user wajib diisi",
		},
		"Title": {
			"required": "Judul acara wajib diisi",
		},
		"Date": {
			"required": "Tanggal wajib diisi",
		},
		"Time": {
			"required": "Waktu wajib diisi",
		},
		"Location": {
			"required": "Lokasi wajib diisi",
		},
		"Description": {
			"required": "Deskripsi wajib diisi",
		},
		"PrimaryColor": {
			"required": "Primary color wajib diisi",
		},
		"SecondaryColor": {
			"required": "Secondary color wajib diisi",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}

func ValidateGenerateLinkRequest(data requests.GenerateLinkRequest) map[string]string {
	messages := map[string]map[string]string{
		"IdInvitation": {
			"required": "Id invitation wajib diisi",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}

func ValidateShareSocialMediaRequest(data requests.ShareSocialMediaRequest) map[string]string {
	messages := map[string]map[string]string{
		"IdInvitation": {
			"required": "Id invitation wajib diisi",
		},
		"NamePlatform": {
			"required": "Nama platform wajib diisi",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}

func ValidateGuestViewRequest(data requests.GuestViewRequest) map[string]string {
	messages := map[string]map[string]string{
		"IdInvitationLink": {
			"required": "Id invitation link wajib diisi",
		},
		"IpAddress": {
			"required": "Ip address wajib diisi",
		},
		"UserAgent": {
			"required": "User agent wajib diisi",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}
