package utilities

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())


func ValidateStruct(s interface{}) map[string]string {
	errs := make(map[string]string)
	validate := validator.New()
	
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
		
			switch err.Tag() {
			case "required":
				errs[field] = field + " is required"
			case "min":
				errs[field] = field + " is too short"
			case "gt":
				errs[field] = field + " must be greater than 0"
			default:
				errs[field] = "Invalid value"
			}
		}
	}
	return errs
}