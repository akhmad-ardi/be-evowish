package lib

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

// Fungsi umum validasi dengan pesan kustom
func ValidateWithCustomMessages(data interface{}, messages map[string]map[string]string) map[string]string {
	err := Validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)

	for _, valErr := range err.(validator.ValidationErrors) {
		field := valErr.Field()
		tag := valErr.Tag()

		if fieldMessages, ok := messages[field]; ok {
			if customMsg, found := fieldMessages[tag]; found {
				errors[field] = customMsg
				continue
			}
		}
		// Default message jika tidak ada custom
		errors[field] = fmt.Sprintf("%s tidak valid", field)
	}

	return errors
}
