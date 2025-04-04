package rules

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUuidRule(t *testing.T) {
	tests := []struct {
		name      string
		uuid      uuid.UUID
		fieldname string
		expectErr bool
	}{
		{"Valid UUID", uuid.New(), "customerId", false},
		{"Nil UUID", uuid.Nil, "orderId", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := UuidRule{tt.fieldname}
			err := rule.Validate(tt.uuid)
			if tt.expectErr {
				assert.NotNil(t, err, "expected error but got nil")
			} else {
				assert.Nil(t, err, "expected no error but got one")
			}
		})
	}
}
