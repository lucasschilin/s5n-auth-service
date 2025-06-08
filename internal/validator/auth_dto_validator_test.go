package validator_test

import (
	"testing"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/validator"
)

func TestIsValidAuthSignupRequest(t *testing.T) {
	tests := []struct {
		name           string
		input          *dto.AuthSignupRequest
		expectedVal    bool
		expectedDetail string
	}{
		{
			name: "Valid Request",
			input: &dto.AuthSignupRequest{
				Email:    "user@test.com",
				Password: "12345678",
			},
			expectedVal:    true,
			expectedDetail: "",
		},
		{
			name: "Missing Email",
			input: &dto.AuthSignupRequest{
				Password: "12345678",
			},
			expectedVal:    false,
			expectedDetail: "Email is required.",
		},
		{
			name: "Missing Password",
			input: &dto.AuthSignupRequest{
				Email: "user@test.com",
			},
			expectedVal:    false,
			expectedDetail: "Password is required.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotVal, gotDetail := validator.IsValidAuthSignupRequest(test.input)
			if gotVal != test.expectedVal || gotDetail != test.expectedDetail {
				t.Errorf(
					"IsValidAuthSignupRequest() = (%v, %v), expected (%v, %v)",
					gotVal, gotDetail, test.expectedVal, test.expectedDetail,
				)
			}
		})
	}
}
