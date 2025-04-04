package rules

import (
	"fmt"
	"time"

	"github.com/hbttundar/diabuddy-errors"
)

// DateMustBeInThePastRule validates that a date is in the past.
type DateMustBeInThePastRule struct {
	FieldName string
}

func (r DateMustBeInThePastRule) Validate(value any) errors.ApiErrors {
	// Ensure the value is a valid time.Time
	date, ok := value.(time.Time)
	if !ok {
		return errors.NewApiError(errors.BadRequestErrorType, r.ValidationErrorMessage("must be a valid date"))
	}

	// Check if the date is in the past
	if !date.Before(time.Now()) {
		return errors.NewApiError(errors.BadRequestErrorType, r.ValidationErrorMessage("must be a date in the past"))
	}

	return nil
}

func (r DateMustBeInThePastRule) ValidationErrorMessage(baseMessage string) string {
	if r.FieldName != "" {
		return fmt.Sprintf("field '%s' %s", r.FieldName, baseMessage)
	}
	return fmt.Sprintf("value of the field %s", baseMessage)
}

func (r DateMustBeInThePastRule) Message() string {
	return "value of the field must be a valid date in the past"
}
