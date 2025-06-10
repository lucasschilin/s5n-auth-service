package auth

import (
	"database/sql"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/port"
	"github.com/lucasschilin/s5n-auth-service/internal/repository"
)

type Service interface {
	Signup(req *dto.AuthSignupRequest) (
		*dto.AuthLoginResponse, *dto.DefaultError,
	)
	Login(req *dto.AuthLoginRequest) (
		*dto.AuthLoginResponse, *dto.DefaultError,
	)
	Refresh(req *dto.AuthRefreshRequest) (
		*dto.AuthRefreshResponse, *dto.DefaultError,
	)
	ForgotPassword(req *dto.AuthForgotPasswordRequest) (
		*dto.DefaultMessageResponse, *dto.DefaultError,
	)
	ResetPassword(req *dto.AuthResetPasswordRequest) (
		*dto.DefaultMessageResponse, *dto.DefaultError,
	)
}

type authService struct {
	UsersDB             *sql.DB
	AuthDB              *sql.DB
	UserRepository      repository.UserRepository
	UserEmailRepository repository.UserEmailRepository
	PasswordRepository  repository.PasswordRepository
	JWTPort             port.JWT
	MailerPort          port.Mailer
}

func NewService(
	usersDB *sql.DB,
	authDB *sql.DB,
	userRepo repository.UserRepository,
	userEmailRepo repository.UserEmailRepository,
	passwordRepo repository.PasswordRepository,
	jwtPort port.JWT,
	mailerPort port.Mailer,

) Service {
	return &authService{
		UsersDB:             usersDB,
		AuthDB:              authDB,
		UserRepository:      userRepo,
		UserEmailRepository: userEmailRepo,
		PasswordRepository:  passwordRepo,
		JWTPort:             jwtPort,
		MailerPort:          mailerPort,
	}
}

const MinPasswordLength = 8

// TODO: refatorar o restante do Service para o formato package
