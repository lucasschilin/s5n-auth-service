package authservice

import (
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/validator"
	"github.com/lucasschilin/s5n-auth-service/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Login(l logger.Logger, req *dto.AuthLoginRequest) (
	*dto.AuthLoginResponse, *dto.DefaultError,
) {
	if val, detail := validator.IsValidAuthLoginRequest(req); !val {
		l.Infof("Invalid request. Detail: %s", detail)
		return nil, errorResponse(http.StatusUnprocessableEntity, detail)
	}

	userEmail, err := s.UserEmailRepository.GetByAddress(&req.Email)
	if err != nil {
		l.Error(err, "")
		return nil, errAuthInternalServerError
	}
	if userEmail == nil {
		l.Infof("Email address = '%s' not found", req.Email)
		return nil, errAuthLoginInvalidCredentials
	}

	password, err := s.PasswordRepository.GetByUser(userEmail.User)
	if err != nil {
		l.Error(err, "")
		return nil, errAuthInternalServerError
	}
	if password == nil {
		l.Infof("Password from user with ID = '%s' not found", userEmail.User)
		return nil, errAuthLoginInvalidCredentials
	}

	if err = bcrypt.CompareHashAndPassword(
		[]byte(password.Password),
		[]byte(req.Password),
	); err != nil {
		if err.Error() != bcrypt.ErrMismatchedHashAndPassword.Error() {
			l.Error(err, "")
			return nil, errAuthInternalServerError
		}

		l.Infof("Password = %s is wrong", req.Password)
		return nil, errAuthLoginInvalidCredentials
	}

	accessToken, err := generateAccessToken(s.TokenManager, userEmail.User)
	if err != nil {
		l.Error(err, "")
		return nil, errAuthInternalServerError
	}

	refreshToken, err := generateRefreshToken(s.TokenManager, userEmail.User)
	if err != nil {
		l.Error(err, "")
		return nil, errAuthInternalServerError
	}

	user, err := s.UserRepository.GetByID(&userEmail.User)
	if err != nil {
		l.Error(err, "")
		return nil, errAuthInternalServerError
	}

	l.Infof("User with ID = '%s' and username = '%s' logged in", user.ID, user.Username)
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
