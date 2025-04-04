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
		return errors.NewApiError(errors.BadRequestErrorType, r.ValidationErrorMessage("is required but got nil"))
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.String:
		if val.Len() == 0 {
			return errors.NewApiError(errors.BadRequestErrorType, r.ValidationErrorMessage("is required but got an empty string"))
		}
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		if val.Len() == 0 {
			return errors.NewApiError(errors.BadRequestErrorType, r.ValidationErrorMessage("is required but got an empty collection"))
		}
	case reflect.Ptr, reflect.Interface:
		if val.IsNil() {
			return errors.NewApiError(errors.BadRequestErrorType, r.ValidationErrorMessage("is required but got a nil pointer or interface"))
		}
	default:
		if fmt.Sprintf("%v", val.Interface()) == "" {
			return errors.NewApiError(errors.BadRequestErrorType, r.ValidationErrorMessage("is required but could not validate the value"))
		}
	}

	return nil
}

func (r RequiredRule) ValidationErrorMessage(baseMessage string) string {
	if r.FieldName != "" {
		return fmt.Sprintf("field '%s' %s", r.FieldName, baseMessage)
	}
	return fmt.Sprintf("value of the field %s", baseMessage)
}
func (r RequiredRule) Message() string {

	return "value for the field is required"
}
