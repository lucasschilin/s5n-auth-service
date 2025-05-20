package service

import (
	"fmt"

	"github.com/lucasschilin/schily-users-api/internal/dto"
	"github.com/lucasschilin/schily-users-api/internal/repository"
)

type AuthService interface {
	Signup(req *dto.AuthSignupRequest) (*dto.AuthSignupResponse, *dto.ControllerError)
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

func (s *authService) Signup(req *dto.AuthSignupRequest) (*dto.AuthSignupResponse, *dto.ControllerError) {
	fmt.Println("Chegou aqui:", req)

	return &dto.AuthSignupResponse{
		User: dto.AuthSignupUserResponse{
			ID:       "tsaufsb",
			Username: "lucaslash",
		},
		AccessToken:  "accesstoketeste",
		RefreshToken: "refreshoketeste",
	}, nil

	// if 1 == 1 {
	// 	return nil, &dto.ControllerError{
	// 		Code:   http.StatusBadRequest,
	// 		Detail: "Bad Request",
	// 	}
	// }

	// TODO: Continue implement de signup feature
}
