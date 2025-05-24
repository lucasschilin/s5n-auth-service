package service

import (
	"net/http"

	"github.com/lucasschilin/schily-users-api/internal/dto"
)

func errorResponse(code int, detail string) *dto.DefaultError {
	return &dto.DefaultError{
		Code:   code,
		Detail: detail,
	}
}

var (
	errAuthSignupInternalServerError = errorResponse(http.StatusInternalServerError, "An error occurred.")
)
