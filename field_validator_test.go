package validation

import (
	"github.com/hbttundar/diabuddy-validation/rules"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldValidatorWithOneRule(t *testing.T) {
	v := NewValidator()

	fv := v.ForField("email", "")
	assert.NotNil(t, fv, "expected FieldValidator but got nil")

	fv.AddRule(rules.EmailRule{})
	assert.NotEmpty(t, v.fields["email"], "expected rules to be added for the email field")
}

func TestFieldValidatorWithTwoRule(t *testing.T) {
	v := NewValidator()

	fv := v.ForField("id", nil)
	assert.NotNil(t, fv, "expected FieldValidator but got nil")

	fv.AddRule(rules.UuidRule{})
	fv.AddRule(rules.RequiredRule{})
	assert.NotEmpty(t, v.fields["id"], "expected rules to be added for the id field")
	assert.Len(t, v.fields["id"], 2, "expected 2 rules for the 'id' field")
}
