package util

import (
	"fmt"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

func ValidateParams(p interface{}) error {
	validate := validator.New()

	s := validate.Struct(p)

	return extractValidationErrors(s)
}

func extractValidationErrors(err error) error {
	if err != nil {
		var errorText []string
		for _, err := range err.(validator.ValidationErrors) {
			errorText = append(errorText, validationErrorToText(err))
		}
		return fmt.Errorf("\nParameter error: %v\n", strings.Join(errorText, "\n"))
	}

	return nil
}

func validationErrorToText(e validator.FieldError) string {
	f := e.Field()
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", f)
	case "max":
		return fmt.Sprintf("%s cannot be greater than %s", f, e.Param())
	case "min":
		return fmt.Sprintf("%s must be greater than %s", f, e.Param())
	}
	return fmt.Sprintf("%s is not valid %s", e.Field(), e.Value())
}
