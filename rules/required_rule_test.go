package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredRule(t *testing.T) {
	rule := RequiredRule{}

	tests := []struct {
		name      string
		input     any
		expectErr bool
	}{
		{"Valid string", "hello", false},
		{"Empty string", "", true},
		{"Nil value", nil, true},
		{"Valid list", []int{1, 2, 3}, false},
		{"Empty list", []int{}, true},
	}

	for _, tt := range tests {
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
