package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailRule(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		fieldName string
		expectErr bool
	}{
		{"basic", "good@email.de", "email", false},
		{"uppercase", "GOOD@EMAIL.DE", "email", false},
		{"umlauts", "gÖöd@email.de", "email", false},
		{"spaces in quotes", `"very good"@email.de`, "email", false},
		{"long with plus", "testing+maurice.birkenfeld__dorfgemeinschaft-froendenberg-west.de@kollex.de", "email", false},
		{"subdomain", "user@sub.domain.com", "email", false},
		{"numeric domain", "user@123.com", "email", false},
		{"local part with dot", "first.last@domain.com", "email", false},
		{"dashes in domain", "user@my-domain.com", "email", false},
		{"valid with subdomain", "user@mail.domain.com", "email", false},
		{"valid numeric domain", "user@123.123.123.123", "email", false},
		{"valid mixed case", "User@Domain.COM", "email", false},
		{"empty string", "", "email", true},
		{"not an email", "not-email", "email", true},
		{"two at signs", "not@good@email.de", "email", true},
		{"spaces without quotes", "not good@email.de", "email", true},
		{"no domain", "no@domain", "email", true},
		{"no user", "@domain.de", "email", true},
		{"with commas", ",,invalidEmail@domain.de", "email", true},
		{"invalid TLD", "user@domain.c", "email", true},
		{"double dots in domain", "user@domain..com", "email", true},
		{"missing at sign", "user.domain.com", "email", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := EmailRule{FieldName: tt.fieldName}
			err := rule.Validate(tt.email)
			if tt.expectErr {
				assert.NotNil(t, err, "expected error but got nil")
			} else {
				assert.Nil(t, err, "expected no error but got one")
			}
		})
	}
}
