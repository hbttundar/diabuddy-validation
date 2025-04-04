package rules

import (
	"fmt"
	"github.com/hbttundar/diabuddy-errors"
	"github.com/nyaruka/phonenumbers"
)

// PhoneNumberRule validates phone numbers.
type PhoneNumberRule struct {
	FieldName string
}

func (p PhoneNumberRule) Validate(value any) errors.ApiErrors {
	phoneNumber, ok := value.(string)
	if !ok {
		return errors.NewApiError(errors.BadRequestErrorType, p.ValidationErrorMessage("has invalid phone number format"))
	}
	parsedNumber, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		return errors.NewApiError(errors.BadRequestErrorType, p.ValidationErrorMessage("is not a valid phone number"), errors.WithInternalError(err))
	}
	if !phonenumbers.IsValidNumber(parsedNumber) {
		return errors.NewApiError(errors.BadRequestErrorType, p.ValidationErrorMessage("is not a valid phone number"))
	}
	return nil
}

func (p PhoneNumberRule) ValidationErrorMessage(baseMessage string) string {
	if p.FieldName != "" {
		return fmt.Sprintf("field '%s' %s", p.FieldName, baseMessage)
	}
	return fmt.Sprintf("value of the field %s", baseMessage)
}

func (p PhoneNumberRule) Message() string {
	return "must be a valid phone number"
}
