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

func IsValidAuthForgotPasswordRequest(req *dto.AuthForgotPasswordRequest) (bool, string) {
	if req.Email == "" {
		return false, "Email is required."
	}

	if req.RedirectUrl == "" {
		return false, "Redirect URL is required."
	}

	return true, ""
}

func IsValidAuthResetPasswordRequest(req *dto.AuthResetPasswordRequest) (bool, string) {
	if req.ResetToken == "" {
		return false, "Reset token is required."
	}

	if req.NewPassword == "" {
		return false, "New password is required."
	}

	return true, ""
}
