package lib

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func getJSONFieldName(data interface{}, field string) string {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		if structField.Name == field {
			jsonTag := structField.Tag.Get("json")
			jsonName := strings.Split(jsonTag, ",")[0] // ambil nama json sebelum koma
			if jsonName != "" && jsonName != "-" {
				return jsonName
			}
		}
	}
	return field // fallback ke nama field asli
}

func ValidateWithCustomMessages(data interface{}, messages map[string]map[string]string) map[string]string {
	err := Validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)

	for _, valErr := range err.(validator.ValidationErrors) {
		field := getJSONFieldName(data, valErr.Field()) // ambil nama dari tag json
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
