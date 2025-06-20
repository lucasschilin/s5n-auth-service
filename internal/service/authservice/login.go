package authservice

import (
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Login(req *dto.AuthLoginRequest) (
	*dto.AuthLoginResponse, *dto.DefaultError,
) {
	if val, detail := validator.IsValidAuthLoginRequest(req); !val {
		return nil, errorResponse(http.StatusUnprocessableEntity, detail)
	}

	userEmail, err := s.UserEmailRepository.GetByAddress(&req.Email)
	if err != nil {
		return nil, errAuthInternalServerError
	}
	if userEmail == nil {
		return nil, errAuthLoginInvalidCredentials
	}

	password, err := s.PasswordRepository.GetByUser(userEmail.User)
	if err != nil {
		return nil, errAuthInternalServerError
	}
	if password == nil {
		return nil, errAuthLoginInvalidCredentials
	}

	if err = bcrypt.CompareHashAndPassword(
		[]byte(password.Password),
		[]byte(req.Password),
	); err != nil {
		return nil, errAuthLoginInvalidCredentials
	}

	accessToken, err := generateAccessToken(s.TokenManager, userEmail.User)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	refreshToken, err := generateRefreshToken(s.TokenManager, userEmail.User)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	user, err := s.UserRepository.GetByID(&userEmail.User)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	return &dto.AuthLoginResponse{
		User: struct {
			ID       string `json:"id"`
			Username string `json:"username"`
		}{
			ID:       user.ID,
			Username: user.Username,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
