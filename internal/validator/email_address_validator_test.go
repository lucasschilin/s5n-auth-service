package validator_test

import (
	"testing"

	"github.com/lucasschilin/s5n-auth-service/internal/validator"
)

func TestIsValidEmailAddress(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected bool
	}{
		{
			name:     "Valid Email",
			input:    []string{"user@gmail.com", "user.user@yahoo.com"},
			expected: true,
		},
		{
			name: "Invalid Length Email",
			input: []string{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@dominio.com",
				"teste@domiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiinio.com",
			},
			expected: false,
		},
		{
			name: "Invalid Regex Email",
			input: []string{
				"lucas@",
				"@email.com",
				"lucasemail.com",
				"lucas@empresa",
				"lucas@@email.com",
				"lucas@em ail.com",
			},
			expected: false,
		},
		{
			name: "Invalid Domain Email",
			input: []string{
				"lucas@dadmmd.xyz",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, address := range test.input {
				got := validator.IsValidEmailAddress(address)
				if got != test.expected {
					t.Errorf(
						"IsValidEmailAddress() = (%v), expected (%v)",
						got, test.expected,
					)
				}
			}
		})
	}
}
