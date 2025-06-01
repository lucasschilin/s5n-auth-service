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

func IsValidAuthLoginRequest(req *dto.AuthLoginRequest) (bool, string) {
	if req.Email == "" {
		return false, "Email is required."
	}

	if req.Password == "" {
		return false, "Password is required."
	}

	return true, ""
}

func IsValidAuthRefreshRequest(req *dto.AuthRefreshRequest) (bool, string) {
	if req.RefreshToken == "" {
		return false, "Refresh token is required."
	}

	return true, ""
}
