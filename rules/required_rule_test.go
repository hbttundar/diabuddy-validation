package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredRule(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		fieldname string
		expectErr bool
	}{
		{"Valid string", "hello", "comment", false},
		{"Empty string", "", "lastname", true},
		{"Nil value", nil, "firstname", true},
		{"Valid list", []int{1, 2, 3}, "orders", false},
		{"Empty list", []int{}, "customers", true},
	}

	for _, tt := range tests {
		rule := RequiredRule{FieldName: tt.fieldname}
		t.Run(tt.name, func(t *testing.T) {
			err := rule.Validate(tt.input)
			if tt.expectErr {
				assert.NotNil(t, err, "expected error but got none")
			} else {
				assert.Nil(t, err, "expected no error but got one")
			}
		})
	}
}
