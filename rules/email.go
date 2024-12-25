package rules

import (
	"github.com/hbttundar/diabuddy-errors"
	"net/mail"
	"regexp"
	"strings"
)

var emailWithDomainRegexp = regexp.MustCompile(`^.+@.+\.[a-z1-9.]{2,}$`)

// EmailRule validates email format.
type EmailRule struct{}

func (e EmailRule) Validate(value any) errors.ApiErrors {
	email, ok := value.(string)
	if !ok {
		return errors.NewApiError(errors.BadRequestErrorType, "invalid email format")
	}
	email = strings.ToLower(email)

	// Parse email using the standard library
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.NewApiError(errors.BadRequestErrorType, "email is not valid", errors.WithInternalError(err))
	}

	// Check for domain part
	if !emailWithDomainRegexp.MatchString(email) {
		return errors.NewApiError(errors.BadRequestErrorType, "email must contain a valid domain")
	}
	return nil
}

func (e EmailRule) Message() string {
	return "must be a valid email address"
}
