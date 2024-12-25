package validation

import "github.com/hbttundar/diabuddy-errors"

// ValidationError represents a single validation error.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// AggregateValidationErrors collects multiple validation errors into a slice.
func AggregateValidationErrors(errors []errors.ApiErrors) []ValidationError {
	var validationErrors []ValidationError
	for _, err := range errors {
		validationErrors = append(validationErrors, ValidationError{
			Field:   "", // Can be populated dynamically in Validator
			Message: err.Error(),
		})
	}
	return validationErrors
}
