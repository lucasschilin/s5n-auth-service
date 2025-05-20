package validator

import (
	"net/http"

	"github.com/lucasschilin/schily-users-api/internal/dto"
)

func IsValidAuthSignupRequest(req *dto.AuthSignupRequest) (bool, *dto.DefaultError) {
	if req.Email == "" {
		return false, &dto.DefaultError{
			Code:   http.StatusBadRequest,
			Detail: "E-mail is required.",
		}
	}

	if req.Password == "" {
		return false, &dto.DefaultError{
			Code:   http.StatusBadRequest,
			Detail: "Password is required.",
		}
	}

	if req.ConfirmPassword == "" {
		return false, &dto.DefaultError{
			Code:   http.StatusBadRequest,
			Detail: "Confirmation password is required.",
		}
	}

	return true, nil
}
