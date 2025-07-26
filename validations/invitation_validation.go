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
