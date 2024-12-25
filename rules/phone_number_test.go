package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhoneNumberRule(t *testing.T) {
	tests := []struct {
		name        string
		phoneNumber string
		expectErr   bool
	}{
		{"Valid phone number US", "+14155552671", false},
		{"Valid phone number DE", "+4917631234567", false},
		{"Invalid phone number without country code", "123456789", true},
		{"Invalid phone number - missing digits", "+1", true},
		{"Invalid phone number with letters", "+1415abcd5671", true},
		{"Valid phone number UK", "+447911123456", false},
		{"Empty phone number", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := PhoneNumberRule{}
			err := rule.Validate(tt.phoneNumber)
			if tt.expectErr {
				assert.NotNil(t, err, "expected error but got nil")
			} else {
				assert.Nil(t, err, "expected no error but got one")
			}
		})
	}
}
