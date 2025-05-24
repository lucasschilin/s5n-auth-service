package validator

import (
	"github.com/lucasschilin/schily-users-api/internal/dto"
)

func IsValidAuthSignupRequest(req *dto.AuthSignupRequest) (bool, string) {
	if req.Email == "" {
		return false, "Email is required."
	}

	if req.Password == "" {
		return false, "Password is required."
	}

	if req.ConfirmPassword == "" {
		return false, "Confirmation password is required."
	}

	return true, ""
}
