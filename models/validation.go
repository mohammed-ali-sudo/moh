package models

import "github.com/go-playground/validator/v10"

// Global validator with required fields enabled on structs.
var validate = validator.New(validator.WithRequiredStructEnabled())

// FirstError returns a simple, general message for the first validation error.
func FirstError(err error) (string, bool) {
	if err == nil {
		return "", true
	}
	if verrs, ok := err.(validator.ValidationErrors); ok && len(verrs) > 0 {
		fe := verrs[0]
		return fe.Field() + " is invalid", false
	}
	return "invalid input", false
}
