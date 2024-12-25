package rules

import (
	"fmt"
	"github.com/hbttundar/diabuddy-errors"
	"reflect"
)

// RequiredRule validates that a value is not empty or nil.
type RequiredRule struct {
	FieldName string
}

func (r RequiredRule) Validate(value any) errors.ApiErrors {
	if value == nil {
		return errors.NewApiError(errors.BadRequestErrorType, fmt.Sprintf("field '%s' is required but got nil", r.FieldName))
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.String:
		if val.Len() == 0 {
			return errors.NewApiError(errors.BadRequestErrorType, fmt.Sprintf("field '%s' is required but got an empty string", r.FieldName))
		}
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		if val.Len() == 0 {
			return errors.NewApiError(errors.BadRequestErrorType, fmt.Sprintf("field '%s' is required but got an empty collection", r.FieldName))
		}
	case reflect.Ptr, reflect.Interface:
		if val.IsNil() {
			return errors.NewApiError(errors.BadRequestErrorType, fmt.Sprintf("field '%s' is required but got a nil pointer or interface", r.FieldName))
		}
	default:
		if fmt.Sprintf("%v", val.Interface()) == "" {
			return errors.NewApiError(errors.BadRequestErrorType, fmt.Sprintf("field '%s' is required but could not validate the value", r.FieldName))
		}
	}

	return nil
}

func (r RequiredRule) Message() string {
	return fmt.Sprintf("field '%s' is required", r.FieldName)
}
