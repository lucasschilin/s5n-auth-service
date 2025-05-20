package service

import (
	"fmt"
	"net/http"

	"github.com/lucasschilin/schily-users-api/internal/dto"
	"github.com/lucasschilin/schily-users-api/internal/repository"
	"github.com/lucasschilin/schily-users-api/internal/validator"
)

type AuthService interface {
	Signup(req *dto.AuthSignupRequest) (*dto.AuthSignupResponse, *dto.DefaultError)
}

type authService struct {
	UserRepository      repository.UserRepository
	UserEmailRepository repository.UserEmailRepository
	PasswordRepository  repository.PasswordRepository
}

func NewAuthService(
	userRepo repository.UserRepository,
	userEmailRepo repository.UserEmailRepository,
	passwordRepo repository.PasswordRepository,
) AuthService {
	return &authService{
		UserRepository:      userRepo,
		UserEmailRepository: userEmailRepo,
		PasswordRepository:  passwordRepo,
	}
}

func (s *authService) Signup(req *dto.AuthSignupRequest) (*dto.AuthSignupResponse, *dto.DefaultError) {
	if _, err := validator.IsValidAuthSignupRequest(req); err != nil {
		return nil, err
	}

	const MinPasswordLength = 8
	if len(req.Password) < MinPasswordLength {
		return nil, &dto.DefaultError{
			Code:   http.StatusBadRequest,
			Detail: fmt.Sprintf("Password must have at least %v characters.", MinPasswordLength),
		}
	}

	if req.Password != req.ConfirmPassword {
		return nil, &dto.DefaultError{
			Code:   http.StatusBadRequest,
			Detail: "Password and confirmation password must match.",
		}
	}

	if !validator.IsValidEmailAddress(req.Email) {
		return nil, &dto.DefaultError{
			Code:   http.StatusBadRequest,
			Detail: "E-mail must be a valid address",
		}
	}

	// TODO: Continue implement de signup feature

	return &dto.AuthSignupResponse{
		User: dto.AuthSignupUserResponse{
			ID:       "tsaufsb",
			Username: "lucaslash",
		},
		AccessToken:  "accesstoketeste",
		RefreshToken: "refreshoketeste",
	}, nil

}
