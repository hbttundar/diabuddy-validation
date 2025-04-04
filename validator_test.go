package validation

import (
	"github.com/google/uuid"
	"github.com/hbttundar/diabuddy-validation/rules"
	"testing"

	_ "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidatorWithInvalidData(t *testing.T) {
	v := NewValidator()

	v.ForField("uuid", uuid.Nil).AddRule(rules.UuidRule{FieldName: "uuid"})
	v.ForField("password", "").AddRule(rules.PasswordRule{FieldName: "password"})
	v.ForField("email", "").AddRule(rules.EmailRule{FieldName: "email"})

	v.Validate()

	assert.True(t, v.HasErrors(), "expected validation errors but got none")

	errors := v.Errors()
	assert.Equal(t, 3, len(errors), "expected 3 errors but got a different count")
	assert.Contains(t, errors[0].Field, "uuid", "expected error for UUID field")
	assert.Contains(t, errors[1].Field, "password", "expected error for password field")
	assert.Contains(t, errors[2].Field, "email", "expected error for email field")

}

func TestValidatorWithValidData(t *testing.T) {
	v := NewValidator()

	v.ForField("uuid", uuid.New()).AddRule(rules.UuidRule{FieldName: "uuid"})
	v.ForField("password", "Password@12345_36").AddRule(rules.PasswordRule{FieldName: "password"})
	v.ForField("email", "hbttundar@gmail.com").AddRule(rules.EmailRule{FieldName: "email"})

	v.Validate()

	assert.Empty(t, v.HasErrors(), "expected no validation errors.")
}
