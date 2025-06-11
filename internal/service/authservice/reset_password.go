package authservice

import (
	"fmt"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) ResetPassword(req *dto.AuthResetPasswordRequest) (
	*dto.DefaultMessageResponse, *dto.DefaultError,
) {
	if val, detail := validator.IsValidAuthResetPasswordRequest(req); !val {
		return nil, errorResponse(http.StatusUnprocessableEntity, detail)
	}

	resetTokenClaims, err := s.JWTPort.ValidateToken(req.Token)
	if err != nil {
		return nil, errAuthInvalidToken
	}

	typeClaim, exists := resetTokenClaims["type"]
	if !exists {
		return nil, errAuthInvalidToken
	}
	typeClaim, ok := typeClaim.(string)
	if !ok {
		return nil, errAuthInvalidToken
	}
	if typeClaim != "reset_password" {
		return nil, errAuthInvalidToken
	}

	if len(req.NewPassword) < MinPasswordLength {
		return nil, errorResponse(
			http.StatusUnprocessableEntity,
			fmt.Sprintf(
				"Password must have at least %v characters.",
				MinPasswordLength,
			),
		)
	}

	subClaim, exists := resetTokenClaims["sub"]
	if !exists {
		return nil, errAuthInvalidToken
	}
	userID, ok := subClaim.(string)
	if !ok {
		return nil, errAuthInvalidToken
	}

	user, err := s.UserRepository.GetByID(&userID)
	if err != nil {
		return nil, errAuthInvalidToken
	}

	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.MinCost)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	if err = s.PasswordRepository.UpdateByUser(
		user.ID, string(bcryptedPassword),
	); err != nil {
		return nil, errAuthInternalServerError
	}

	return &dto.DefaultMessageResponse{
		Message: "Password changed successfully.",
	}, nil
}
