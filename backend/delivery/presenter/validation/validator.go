package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Init validator dan register custom rules
func init() {
	validate = validator.New()

	// Register custom validation untuk "not_blank"
	validate.RegisterValidation("not_blank", func(fl validator.FieldLevel) bool {
		return strings.TrimSpace(fl.Field().String()) != ""
	})
}

// ValidateStruct memvalidasi struct apapun
func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}
