package rules

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hbttundar/diabuddy-errors"
)

// UuidRule validates UUIDs.
type UuidRule struct {
	FieldName string
}

func (u UuidRule) Validate(value any) errors.ApiErrors {
	id, ok := value.(uuid.UUID)
	if !ok {
		return errors.NewApiError(errors.BadRequestErrorType, u.ValidationErrorMessage("has invalid UUID format"))
	}
	if id == uuid.Nil {
		return errors.NewApiError(errors.BadRequestErrorType, u.ValidationErrorMessage("cannot be nil as a UUID"))
	}
	return nil
}

func (u UuidRule) ValidationErrorMessage(baseMessage string) string {
	if u.FieldName != "" {
		return fmt.Sprintf("field '%s' %s", u.FieldName, baseMessage)
	}
	return fmt.Sprintf("value of the field %s", baseMessage)
}

func (u UuidRule) Message() string {
	return "must be a valid UUID"
}
