package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordRule(t *testing.T) {
	tests := []struct {
		name      string
		password  string
		fieldname string
		expectErr bool
	}{
		{"Valid password with uppercase, lowercase, number, and special character", "ValidPassword123!", "password", false},
		{"Valid password with minimum requirements", "Passw0rd!", "password", false},
		{"Empty password", "", "password", true},
		{"Password too short", "Pass1!", "password", true},
		{"Password missing number", "Password!", "password", true},
		{"Password missing uppercase letter", "password123!", "password", true},
		{"Password missing lowercase letter", "PASSWORD123!", "password", true},
		{"Password missing special character", "Password123", "password", true},
		{"Password with only numbers", "12345678!", "password", true},
		{"Password with only letters", "PasswordWithoutNumbers!", "password", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := PasswordRule{FieldName: tt.fieldname}
			err := rule.Validate(tt.password)
			if tt.expectErr {
				assert.NotNil(t, err, "expected error but got nil")
			} else {
				assert.Nil(t, err, "expected no error but got one")
			}
		})
	}
}
