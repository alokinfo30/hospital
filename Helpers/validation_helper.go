// validation_helper.go

package Helper

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct based on the tags defined in the struct
func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}
