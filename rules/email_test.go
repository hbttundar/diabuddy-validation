package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailRule(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		expectErr bool
	}{
		{"basic", "good@email.de", false},
		{"uppercase", "GOOD@EMAIL.DE", false},
		{"umlauts", "gÖöd@email.de", false},
		{"spaces in quotes", `"very good"@email.de`, false},
		{"long with plus", "testing+maurice.birkenfeld__dorfgemeinschaft-froendenberg-west.de@kollex.de", false},
		{"subdomain", "user@sub.domain.com", false},
		{"numeric domain", "user@123.com", false},
		{"local part with dot", "first.last@domain.com", false},
		{"dashes in domain", "user@my-domain.com", false},
		{"valid with subdomain", "user@mail.domain.com", false},
		{"valid numeric domain", "user@123.123.123.123", false},
		{"valid mixed case", "User@Domain.COM", false},
		{"empty string", "", true},
		{"not an email", "not-email", true},
		{"two at signs", "not@good@email.de", true},
		{"spaces without quotes", "not good@email.de", true},
		{"no domain", "no@domain", true},
		{"no user", "@domain.de", true},
		{"with commas", ",,invalidEmail@domain.de", true},
		{"invalid TLD", "user@domain.c", true},
		{"double dots in domain", "user@domain..com", true},
		{"missing at sign", "user.domain.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := EmailRule{}
			err := rule.Validate(tt.email)
			if tt.expectErr {
				assert.NotNil(t, err, "expected error but got nil")
			} else {
				assert.Nil(t, err, "expected no error but got one")
			}
		})
	}
}
