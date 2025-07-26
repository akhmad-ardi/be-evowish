package validations

import (
	"be-undangan-digital/lib"
	"be-undangan-digital/requests"
)

// --- Validasi Register ---
func ValidateRegisterRequest(data requests.RegisterRequest) map[string]string {
	messages := map[string]map[string]string{
		"Name": {
			"required": "Nama wajib diisi",
			"min":      "Nama minimal 2 karakter",
		},
		"Email": {
			"required": "Email wajib diisi",
			"email":    "Format email tidak valid",
		},
		"Password": {
			"required": "Password wajib diisi",
			"min":      "Password minimal 6 karakter",
		},
		"ConfirmPassword": {
			"required": "Konfirmasi password wajib diisi",
			"eqfield":  "Konfirmasi password harus sama dengan password",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}

// --- Validasi Login ---
func ValidateLoginRequest(data requests.LoginRequest) map[string]string {
	messages := map[string]map[string]string{
		"Email": {
			"required": "Email wajib diisi",
			"email":    "Format email tidak valid",
		},
		"Password": {
			"required": "Password wajib diisi",
			"min":      "Password minimal 6 karakter",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}
