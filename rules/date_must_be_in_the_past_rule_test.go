package rules

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateMustBeInThePastRule(t *testing.T) {
	rule := DateMustBeInThePastRule{FieldName: "birth_date"}

	tests := []struct {
		name      string
		input     any
		expectErr bool
	}{
		{"Valid past date", time.Now().Add(-24 * time.Hour), false},
		{"Future date", time.Now().Add(24 * time.Hour), true},
		{"Invalid type", "not-a-date", true},
		{"Nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := rule.Validate(tt.input)
			if tt.expectErr {
				assert.NotNil(t, err, "expected error but got nil")
			} else {
				assert.Nil(t, err, "expected no error but got one")
			}
		})
	}
}
