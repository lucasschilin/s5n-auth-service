package authservice

import (
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/validator"
)

func (s *authService) Refresh(req *dto.AuthRefreshRequest) (
	*dto.AuthRefreshResponse, *dto.DefaultError,
) {
	if val, detail := validator.IsValidAuthRefreshRequest(req); !val {
		return nil, errorResponse(http.StatusUnprocessableEntity, detail)
	}

	refreshTokenClaims, err := s.JWTPort.ValidateToken(req.RefreshToken)
	if err != nil {
		return nil, errAuthInvalidToken
	}

	typeClaim, exists := refreshTokenClaims["type"]
	if !exists {
		return nil, errAuthInvalidToken
	}
	if typeClaim, ok := typeClaim.(string); !ok || typeClaim != "refresh_token" {
		return nil, errAuthInvalidToken
	}

	sub, exists := refreshTokenClaims["sub"]
	if !exists {
		return nil, errAuthInvalidToken
	}
	userID, ok := sub.(string)
	if !ok {
		return nil, errAuthInvalidToken
	}

	user, _ := s.UserRepository.GetByID(&userID)
	if user == nil {
		return nil, errAuthInvalidToken
	}

	accessToken, err := generateAccessToken(s.JWTPort, user.ID)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	return &dto.AuthRefreshResponse{
		AccessToken: accessToken,
	}, nil
}
