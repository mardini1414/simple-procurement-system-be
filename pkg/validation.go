package pkg

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(data any) map[string]string {
	var errors = make(map[string]string)

	if err := validate.Struct(data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			print(err)
			errors[ToSnakeCase(err.Field())] = validationMessage(err)
		}
	}

	return errors
}

func validationMessage(err validator.FieldError) string {
	p := err.Param()
	switch err.Tag() {
	case "required":
		return "field ini wajib diisi"
	case "email":
		return "format email tidak valid"
	case "min":
		return fmt.Sprintf("minimal %s karakater", p)
	case "max":
		return fmt.Sprintf("maximal %s karakater", p)
	case "gte":
		return fmt.Sprintf("harus lebih besar atau sama dengan %s", p)
	case "lte":
		return fmt.Sprintf("harus lebih kecil atau sama dengan %s", p)
	case "number":
		return "harus berupa number"
	default:
		return "field tidak valid"
	}
}
