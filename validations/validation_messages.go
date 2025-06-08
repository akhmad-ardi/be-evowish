package validations

import (
	"github.com/go-playground/validator/v10"
)

func ValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " harus diisi"
	case "email":
		return "Email harus dalam format yang valid"
	case "min":
		return fe.Field() + " minimal " + fe.Param() + " karakter"
	case "eqfield":
		return fe.Field() + " harus sama dengan " + fe.Param()
	default:
		return fe.Field() + " tidak valid"
	}
}
