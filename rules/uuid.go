package rules

import (
	"github.com/google/uuid"
	"github.com/hbttundar/diabuddy-errors"
)

// UuidRule validates UUIDs.
type UuidRule struct{}

func (u UuidRule) Validate(value any) errors.ApiErrors {
	id, ok := value.(uuid.UUID)
	if !ok {
		return errors.NewApiError(errors.BadRequestErrorType, "invalid UUID format")
	}
	if id == uuid.Nil {
		return errors.NewApiError(errors.BadRequestErrorType, "UUID must not be nil")
	}
	return nil
}

func (u UuidRule) Message() string {
	return "must be a valid UUID"
}
