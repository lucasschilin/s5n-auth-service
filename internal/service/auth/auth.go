package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/port"
	"github.com/lucasschilin/s5n-auth-service/internal/repository"
	"github.com/lucasschilin/s5n-auth-service/internal/util"
	"github.com/lucasschilin/s5n-auth-service/internal/validator"
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

	accessToken, err := generateAccessToken(s.JWTPort, userEmail.User)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	refreshToken, err := generateRefreshToken(s.JWTPort, userEmail.User)
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

func (s *authService) ForgotPassword(req *dto.AuthForgotPasswordRequest) (
	*dto.DefaultMessageResponse, *dto.DefaultError,
) {
	messageResponse := dto.DefaultMessageResponse{
		Message: "Email sent.",
	}

	if val, detail := validator.IsValidAuthForgotPasswordRequest(req); !val {
		return nil, errorResponse(http.StatusUnprocessableEntity, detail)
	}

	if !validator.IsValidEmailAddress(req.Email) {
		return nil, errorResponse(
			http.StatusUnprocessableEntity, "Email must be a valid address",
		)
	}

	allowedRedirectHosts := []string{"s5n.com.br"}

	redirectUrlWithoutHttp := strings.ReplaceAll(strings.ToLower(req.RedirectUrl), "http://", "")
	redirectUrlWithoutHttp = strings.ReplaceAll(redirectUrlWithoutHttp, "https://", "")
	redirectUrlHost := strings.Split(redirectUrlWithoutHttp, "/")[0]

	finded, _ := util.InStringSlice(allowedRedirectHosts, redirectUrlHost)
	if !finded {
		return nil, errorResponse(
			http.StatusUnprocessableEntity,
			"Redirect URL must be a valid and allowed URL",
		)
	}

	userEmail, err := s.UserEmailRepository.GetByAddress(&req.Email)
	if err != nil {
		return nil, errAuthInternalServerError
	}
	if userEmail == nil {
		return &messageResponse, nil
	}

	user, err := s.UserRepository.GetByID(&userEmail.User)
	if err != nil {
		return nil, errAuthInternalServerError
	}
	if user == nil {
		return &messageResponse, nil
	}

	exp := time.Now().Add(5 * time.Minute).Unix()
	token, err := generateToken(s.JWTPort, "reset_password", int(exp), user.ID)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	link := req.RedirectUrl + "?t=" + token

	subject := "🔐 Recupere o acesso à sua conta – redefina sua senha"
	body := fmt.Sprintf(`<p>Olá %s, como vai?</p>
			<div>
				<p>Para redefinir sua senha e recuperar sua conta, copie e cole este link no seu navegador:</p>
				<p>%s</p>
			</div>`, user.Username, link)

	err = s.MailerPort.NewMessage().
		Subject(&subject).
		Body(&body).
		To(&[]string{userEmail.Address}).
		Send()
	if err != nil {
		fmt.Printf("Erro ao enviar email: %v\n", err)
		return nil, errAuthInternalServerError
	}

	return &messageResponse, nil
}

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
