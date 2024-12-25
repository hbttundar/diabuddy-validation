package validation

import (
	"github.com/hbttundar/diabuddy-validation/rules"
	"testing"

	_ "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	v := NewValidator()

	v.ForField("email").AddRule(rules.EmailRule{})
	v.ForField("password").AddRule(rules.PasswordRule{})
	v.ForField("uuid").AddRule(rules.UuidRule{})

	v.Validate()

	assert.True(t, v.HasErrors(), "expected validation errors but got none")

	errors := v.Errors()
	assert.Equal(t, 3, len(errors), "expected 3 errors but got a different count")
	assert.Contains(t, errors[0].Field, "email", "expected error for email field")
	assert.Contains(t, errors[1].Field, "password", "expected error for password field")
	assert.Contains(t, errors[2].Field, "uuid", "expected error for UUID field")
}
