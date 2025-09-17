package models

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Shared validator instance
var validate = validator.New()

// FirstError returns first validation error in a friendly format.
func FirstError(err error) (string, bool) {
	if err == nil {
		return "", true
	}
	verrs, ok := err.(validator.ValidationErrors)
	if !ok || len(verrs) == 0 {
		return err.Error(), false
	}
	fe := verrs[0]

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field()), false
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", fe.Field(), fe.Param()), false
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", fe.Field(), fe.Param()), false
	case "email":
		return "Invalid email format", false
	case "alpha":
		return fmt.Sprintf("%s must contain only letters (no spaces)", fe.Field()), false
	case "alphanum":
		return fmt.Sprintf("%s must be alphanumeric (no spaces)", fe.Field()), false
	case "alpha_space":
		return fmt.Sprintf("%s must contain only letters and spaces", fe.Field()), false
	case "alphanum_space":
		return fmt.Sprintf("%s must be letters, numbers, and spaces", fe.Field()), false
	case "required_with":
		return fmt.Sprintf("%s is required when %s is present", fe.Field(), fe.Param()), false
	case "required_without":
		return fmt.Sprintf("%s is required when %s is missing", fe.Field()), false
	case "email_domain":
		return "Email must be a @school.edu address", false
	case "uuid", "uuid4", "uuid_opt":
		return fmt.Sprintf("%s must be a valid UUID", fe.Field()), false
	case "json":
		return fmt.Sprintf("%s must be valid JSON", fe.Field()), false
	case "oneof":
		return fmt.Sprintf("%s must be one of %s", fe.Field(), fe.Param()), false
	case "gt", "gte", "lt", "lte":
		return fmt.Sprintf("%s must satisfy %s %s", fe.Field(), fe.Tag(), fe.Param()), false
	case "gt_mfg":
		return "ExpireDate must be after MfgDate", false
	case "gt_valid_from":
		return "ValidTo must be after ValidFrom", false
	default:
		return fmt.Sprintf("%s is invalid", fe.Field()), false
	}
}

// Register custom field and struct validations.
func init() {

	// notblank: trims spaces before checking
	_ = validate.RegisterValidation("notblank", func(fl validator.FieldLevel) bool {
		s, ok := fl.Field().Interface().(string)
		if !ok {
			return true
		}
		return strings.TrimSpace(s) != ""
	})

	// alpha_space: letters + spaces (unicode aware)
	alphaSpace := regexp.MustCompile(`^\p{L}+(?:[ ]\p{L}+)*$`)
	_ = validate.RegisterValidation("alpha_space", func(fl validator.FieldLevel) bool {
		s, ok := fl.Field().Interface().(string)
		if !ok {
			return true
		}
		s = strings.TrimSpace(s)
		if s == "" {
			return false
		}
		return alphaSpace.MatchString(s)
	})

	// alphanum_space: letters/numbers/spaces
	alphanumSpace := regexp.MustCompile(`^[\p{L}\p{N} ]+$`)
	_ = validate.RegisterValidation("alphanum_space", func(fl validator.FieldLevel) bool {
		s, ok := fl.Field().Interface().(string)
		if !ok {
			return true
		}
		s = strings.TrimSpace(s)
		if s == "" {
			return false
		}
		return alphanumSpace.MatchString(s)
	})

	// email_domain: require @school.edu
	_ = validate.RegisterValidation("email_domain", func(fl validator.FieldLevel) bool {
		s, ok := fl.Field().Interface().(string)
		if !ok || s == "" {
			return true
		}
		return strings.HasSuffix(strings.ToLower(strings.TrimSpace(s)), "@school.edu")
	})

	// uuid_opt: allow empty string; else must be UUID (use built-in "uuid")
	_ = validate.RegisterValidation("uuid_opt", func(fl validator.FieldLevel) bool {
		s, ok := fl.Field().Interface().(string)
		if !ok || s == "" {
			return true
		}
		type tmp struct {
			V string `validate:"uuid"`
		}
		return validate.Struct(&tmp{V: s}) == nil
	})

	// json: ensures the field contains valid JSON (supports string, []byte, json.RawMessage)
	_ = validate.RegisterValidation("json", func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface()
		switch v := val.(type) {
		case string:
			if strings.TrimSpace(v) == "" {
				return false
			}
			var tmp any
			return json.Unmarshal([]byte(v), &tmp) == nil
		case []byte:
			var tmp any
			return json.Unmarshal(v, &tmp) == nil
		case json.RawMessage:
			if v == nil {
				return false
			}
			var tmp any
			return json.Unmarshal(v, &tmp) == nil
		default:
			return false
		}
	})
}
