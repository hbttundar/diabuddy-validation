package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordRule(t *testing.T) {
	tests := []struct {
		name      string
		password  string
		expectErr bool
	}{
		{"Valid password with uppercase, lowercase, number, and special character", "ValidPassword123!", false},
		{"Valid password with minimum requirements", "Passw0rd!", false},
		{"Empty password", "", true},
		{"Password too short", "Pass1!", true},
		{"Password missing number", "Password!", true},
		{"Password missing uppercase letter", "password123!", true},
		{"Password missing lowercase letter", "PASSWORD123!", true},
		{"Password missing special character", "Password123", true},
		{"Password with only numbers", "12345678!", true},
		{"Password with only letters", "PasswordWithoutNumbers!", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := PasswordRule{}
			err := rule.Validate(tt.password)
			if tt.expectErr {
				assert.NotNil(t, err, "expected error but got nil")
			} else {
				assert.Nil(t, err, "expected no error but got one")
			}
		})
	}
}
