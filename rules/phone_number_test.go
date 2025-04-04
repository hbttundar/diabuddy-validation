package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhoneNumberRule(t *testing.T) {
	tests := []struct {
		name        string
		phoneNumber string
		fieldName   string
		expectErr   bool
	}{
		{"Valid phone number US", "+14155552671", "mobile", false},
		{"Valid phone number DE", "+4917631234567", "mobile", false},
		{"Invalid phone number without country code", "123456789", "phone", true},
		{"Invalid phone number - missing digits", "+1", "phone", true},
		{"Invalid phone number with letters", "+1415abcd5671", "mobile", true},
		{"Valid phone number UK", "+447911123456", "mobile", false},
		{"Empty phone number", "", "mobile", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := PhoneNumberRule{FieldName: tt.fieldName}
			err := rule.Validate(tt.phoneNumber)
			if tt.expectErr {
				assert.NotNil(t, err, "expected error but got nil")
			} else {
				assert.Nil(t, err, "expected no error but got one")
			}
		})
	}
}
