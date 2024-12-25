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
		return errors.NewApiError(errors.BadRequestErrorType, fmt.Sprintf("field '%s' must be a valid date", r.FieldName))
	}

	// Check if the date is in the past
	if !date.Before(time.Now()) {
		return errors.NewApiError(errors.BadRequestErrorType, fmt.Sprintf("field '%s' must be a date in the past", r.FieldName))
	}

	return nil
}

func (r DateMustBeInThePastRule) Message() string {
	return fmt.Sprintf("field '%s' must be a date in the past", r.FieldName)
}
