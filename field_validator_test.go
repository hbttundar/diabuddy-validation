package validation

import (
	"github.com/hbttundar/diabuddy-validation/rules"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldValidator(t *testing.T) {
	v := NewValidator()

	fv := v.ForField("email")
	assert.NotNil(t, fv, "expected FieldValidator but got nil")

	fv.AddRule(rules.EmailRule{})
	assert.NotEmpty(t, v.fields["email"], "expected rules to be added for the email field")

	fv.AddRule(rules.PasswordRule{})
	assert.Equal(t, 2, len(v.fields["email"]), "expected 2 rules for the email field")
}
