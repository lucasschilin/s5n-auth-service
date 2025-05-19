package service

import (
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
	return nil, nil
}
