package rules

import (
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
type PasswordRule struct{}

func (p PasswordRule) Validate(value any) errors.ApiErrors {
	password, ok := value.(string)
	if !ok {
		return errors.NewApiError(errors.BadRequestErrorType, "invalid password format")
	}
	if len(password) < minPasswordLength {
		return errors.NewApiError(errors.BadRequestErrorType, "password must be at least 8 characters long")
	}
	if !uppercaseRegexp.MatchString(password) {
		return errors.NewApiError(errors.BadRequestErrorType, "password must contain at least one uppercase letter")
	}
	if !lowercaseRegexp.MatchString(password) {
		return errors.NewApiError(errors.BadRequestErrorType, "password must contain at least one lowercase letter")
	}
	if !digitRegexp.MatchString(password) {
		return errors.NewApiError(errors.BadRequestErrorType, "password must contain at least one digit")
	}
	if !specialCharRegexp.MatchString(password) {
		return errors.NewApiError(errors.BadRequestErrorType, "password must contain at least one special character")
	}
	return nil
}

func (p PasswordRule) Message() string {
	return "must be a strong password"
}
