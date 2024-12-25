package rules

import (
	"github.com/hbttundar/diabuddy-errors"
	"github.com/nyaruka/phonenumbers"
)

// PhoneNumberRule validates phone numbers.
type PhoneNumberRule struct{}

func (p PhoneNumberRule) Validate(value any) errors.ApiErrors {
	phoneNumber, ok := value.(string)
	if !ok {
		return errors.NewApiError(errors.BadRequestErrorType, "invalid phone number format")
	}
	parsedNumber, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		return errors.NewApiError(errors.BadRequestErrorType, "phone number is not valid", errors.WithInternalError(err))
	}
	if !phonenumbers.IsValidNumber(parsedNumber) {
		return errors.NewApiError(errors.BadRequestErrorType, "phone number is not valid")
	}
	return nil
}

func (p PhoneNumberRule) Message() string {
	return "must be a valid phone number"
}
