package authservice

import (
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
)

func errorResponse(code int, detail string) *dto.DefaultError {
	return &dto.DefaultError{
		Code:   code,
		Detail: detail,
	}
}

var (
	errAuthInternalServerError     = errorResponse(http.StatusInternalServerError, "An error occurred.")
	errAuthLoginInvalidCredentials = errorResponse(http.StatusUnauthorized, "Invalid credentials.")
	errAuthInvalidToken            = errorResponse(http.StatusUnauthorized, "Invalid token.")
)
