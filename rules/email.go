package rules

import (
	"fmt"
	"github.com/hbttundar/diabuddy-errors"
	"net/mail"
	"regexp"
	"strings"
)

var emailWithDomainRegexp = regexp.MustCompile(`^.+@.+\.[a-z1-9.]{2,}$`)

// EmailRule validates email format.
type EmailRule struct {
	FieldName string
}

func (e EmailRule) Validate(value any) errors.ApiErrors {
	email, ok := value.(string)
	if !ok {
		return errors.NewApiError(errors.BadRequestErrorType, e.ValidationErrorMessage("has invalid email format"))
	}
	email = strings.ToLower(email)

	// Parse email using the standard library
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.NewApiError(errors.BadRequestErrorType, e.ValidationErrorMessage("is not valid email"), errors.WithInternalError(err))
	}

	// Check for domain part
	if !emailWithDomainRegexp.MatchString(email) {
		return errors.NewApiError(errors.BadRequestErrorType, e.ValidationErrorMessage("must contain a valid domain for email"))
	}
	return nil
}

func (e EmailRule) ValidationErrorMessage(baseMessage string) string {
	if e.FieldName != "" {
		return fmt.Sprintf("field '%s' %s", e.FieldName, baseMessage)
	}
	return fmt.Sprintf("value of the field %s", baseMessage)
}

func (e EmailRule) Message() string {
	return "value of the field must be a valid email address"
}
