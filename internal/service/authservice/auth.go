package authservice

import (
	"database/sql"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/integrations/mailer"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/password"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/user"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/useremail"
	"github.com/lucasschilin/s5n-auth-service/internal/service/authservice/jwt"
	"github.com/lucasschilin/s5n-auth-service/pkg/logger"
)

type Service interface {
	Signup(l logger.Logger, req *dto.AuthSignupRequest) (
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
	UserRepository      user.Repository
	UserEmailRepository useremail.Repository
	PasswordRepository  password.Repository
	TokenManager        jwt.TokenManager
	Mailer              mailer.Mailer
}

func NewService(
	usersDB *sql.DB,
	authDB *sql.DB,
	userRepo user.Repository,
	userEmailRepo useremail.Repository,
	passwordRepo password.Repository,
	tokenManager jwt.TokenManager,
	mailerPort mailer.Mailer,

) Service {
	return &authService{
		UsersDB:             usersDB,
		AuthDB:              authDB,
		UserRepository:      userRepo,
		UserEmailRepository: userEmailRepo,
		PasswordRepository:  passwordRepo,
		TokenManager:        tokenManager,
		Mailer:              mailerPort,
	}
}

const MinPasswordLength = 8
