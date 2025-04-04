package rules

import (
	"fmt"
	"github.com/hbttundar/diabuddy-errors"
	"regexp"
)

const minPasswordLength = 8

var (
	uppercaseRegexp   = regexp.MustCompile(`[A-Z]`)
	lowercaseRegexp   = regexp.MustCompile(`[a-z]`)
	digitRegexp       = regexp.MustCompile(`[0-9]`)
	specialCharRegexp = regexp.MustCompile(`[!@#~$%^&*()_+={}\[\]|\\:;"'<>,.?/-]`)
)

// PasswordRule validates password strength.
type PasswordRule struct {
	FieldName string
}

func (p PasswordRule) Validate(value any) errors.ApiErrors {
	password, ok := value.(string)
	if !ok {
		return errors.NewApiError(errors.BadRequestErrorType, p.ValidationErrorMessage("has invalid password format"))
	}
	if len(password) < minPasswordLength {
		return errors.NewApiError(errors.BadRequestErrorType, p.ValidationErrorMessage("must be at least 8 characters long as a password"))
	}
	if !uppercaseRegexp.MatchString(password) {
		return errors.NewApiError(errors.BadRequestErrorType, p.ValidationErrorMessage("must contain at least one uppercase letter as a password"))
	}
	if !lowercaseRegexp.MatchString(password) {
		return errors.NewApiError(errors.BadRequestErrorType, p.ValidationErrorMessage("must contain at least one lowercase letter as a password"))
	}
	if !digitRegexp.MatchString(password) {
		return errors.NewApiError(errors.BadRequestErrorType, p.ValidationErrorMessage("must contain at least one digit as a password"))
	}
	if !specialCharRegexp.MatchString(password) {
		return errors.NewApiError(errors.BadRequestErrorType, p.ValidationErrorMessage("field '%s' must contain at least one special character as a password"))
	}
	return nil
}

func (p PasswordRule) ValidationErrorMessage(baseMessage string) string {
	if p.FieldName != "" {
		return fmt.Sprintf("field '%s' %s", p.FieldName, baseMessage)
	}
	return fmt.Sprintf("value of the field %s", baseMessage)
}

func (p PasswordRule) Message() string {
	return "value of the field must be a valid strong password"
}
