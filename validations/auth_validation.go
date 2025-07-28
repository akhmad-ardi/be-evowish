package validations

import (
	"be-undangan-digital/lib"
	"be-undangan-digital/requests"
)

// --- Validasi Register ---
func ValidateRegisterRequest(data requests.RegisterRequest) map[string]string {
	messages := map[string]map[string]string{
		"name": {
			"required": "Nama wajib diisi",
			"min":      "Nama minimal 2 karakter",
		},
		"email": {
			"required": "Email wajib diisi",
			"email":    "Format email tidak valid",
		},
		"password": {
			"required": "Password wajib diisi",
			"min":      "Password minimal 6 karakter",
		},
		"confirm_password": {
			"required": "Konfirmasi password wajib diisi",
			"eqfield":  "Konfirmasi password harus sama dengan password",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}

// --- Validasi Login ---
func ValidateLoginRequest(data requests.LoginRequest) map[string]string {
	messages := map[string]map[string]string{
		"email": {
			"required": "Email wajib diisi",
			"email":    "Format email tidak valid",
		},
		"password": {
			"required": "Password wajib diisi",
			"min":      "Password minimal 6 karakter",
		},
	}

	return lib.ValidateWithCustomMessages(data, messages)
}
