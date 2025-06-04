package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/aidarkhanov/nanoid"
	"golang.org/x/crypto/bcrypt"

	"github.com/lucasschilin/schily-users-api/internal/dto"
	"github.com/lucasschilin/schily-users-api/internal/port"
	"github.com/lucasschilin/schily-users-api/internal/repository"
	"github.com/lucasschilin/schily-users-api/internal/util"
	"github.com/lucasschilin/schily-users-api/internal/validator"
)

type AuthService interface {
	Signup(req *dto.AuthSignupRequest) (
		*dto.AuthLoginResponse, *dto.DefaultError,
	)
	Login(req *dto.AuthLoginRequest) (
		*dto.AuthLoginResponse, *dto.DefaultError,
	)
	Refresh(req *dto.AuthRefreshRequest) (
		*dto.AuthRefreshResponse, *dto.DefaultError,
	)
	ForgotPassword(req *dto.AuthForgotPasswordRequest) *dto.DefaultError
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

func NewAuthService(
	usersDB *sql.DB,
	authDB *sql.DB,
	userRepo repository.UserRepository,
	userEmailRepo repository.UserEmailRepository,
	passwordRepo repository.PasswordRepository,
	jwtPort port.JWT,
	mailerPort port.Mailer,

) AuthService {
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

func (s *authService) Signup(req *dto.AuthSignupRequest) (
	*dto.AuthLoginResponse, *dto.DefaultError,
) {
	if val, detail := validator.IsValidAuthSignupRequest(req); !val {
		return nil, errorResponse(http.StatusUnprocessableEntity, detail)
	}

	const MinPasswordLength = 8
	if len(req.Password) < MinPasswordLength {
		return nil, errorResponse(
			http.StatusUnprocessableEntity,
			fmt.Sprintf(
				"Password must have at least %v characters.",
				MinPasswordLength,
			),
		)
	}

	if req.Password != req.ConfirmPassword {
		return nil, errorResponse(
			http.StatusUnprocessableEntity,
			"Password and confirmation password must match.",
		)
	}

	if !validator.IsValidEmailAddress(req.Email) {
		return nil, errorResponse(
			http.StatusUnprocessableEntity, "Email must be a valid address",
		)
	}

	userEmail, err := s.UserEmailRepository.GetByAddress(&req.Email)
	if err != nil {
		return nil, errorResponse(
			http.StatusInternalServerError, "Email cannot be validated.",
		)
	}
	if userEmail != nil {
		return nil, errorResponse(http.StatusConflict, "Email already in use.")
	}

	usersTX, err := s.UsersDB.Begin()
	if err != nil {
		return nil, errAuthInternalServerError
	}
	authTX, err := s.AuthDB.Begin()
	if err != nil {
		return nil, errAuthInternalServerError
	}
	defer usersTX.Rollback()
	defer authTX.Rollback()

	emailUsername := strings.Split(req.Email, "@")[0]
	maxUsernameLength := 13
	if len(emailUsername) < maxUsernameLength {
		maxUsernameLength = len(emailUsername)
	}
	username := strings.ToLower(emailUsername[:maxUsernameLength])
	user, err := s.UserRepository.GetByUsername(&username)
	if err != nil {
		return nil, errAuthInternalServerError
	}
	if user != nil {
		sufix, err := nanoid.Generate(nanoid.DefaultAlphabet, 5)
		if err != nil {
			return nil, errAuthInternalServerError
		}
		username = strings.ToLower(fmt.Sprintf("%s_%s", username, sufix))
	}

	newUser, err := s.UserRepository.CreateWithTX(usersTX, &username)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	verifyToken, err := nanoid.Generate(nanoid.DefaultAlphabet, 50)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	_, err = s.UserEmailRepository.CreateWithTX(
		usersTX, &newUser.ID, &req.Email, &verifyToken,
	)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	_, err = s.PasswordRepository.CreateWithTX(
		authTX, newUser.ID, string(bcryptedPassword),
	)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	accessToken, err := generateAccessToken(s.JWTPort, newUser.ID)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	refreshToken, err := generateRefreshToken(s.JWTPort, newUser.ID)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	usersTX.Commit()
	authTX.Commit()

	return &dto.AuthLoginResponse{
		User: struct {
			ID       string `json:"id"`
			Username string `json:"username"`
		}{
			ID:       newUser.ID,
			Username: newUser.Username,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

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

func (s *authService) ForgotPassword(req *dto.AuthForgotPasswordRequest) *dto.DefaultError {
	if val, detail := validator.IsValidAuthForgotPasswordRequest(req); !val {
		return errorResponse(http.StatusUnprocessableEntity, detail)
	}

	if !validator.IsValidEmailAddress(req.Email) {
		return errorResponse(
			http.StatusUnprocessableEntity, "Email must be a valid address",
		)
	}

	allowedRedirectHosts := []string{"s5n.com.br"}

	redirectUrlWithoutHttp := strings.ReplaceAll(strings.ToLower(req.RedirectUrl), "http://", "")
	redirectUrlWithoutHttp = strings.ReplaceAll(redirectUrlWithoutHttp, "https://", "")
	redirectUrlHost := strings.Split(redirectUrlWithoutHttp, "/")[0]

	finded, _ := util.InStringSlice(allowedRedirectHosts, redirectUrlHost)
	if !finded {
		return errorResponse(
			http.StatusUnprocessableEntity,
			"Redirect URL must be a valid and allowed URL",
		)
	}

	userEmail, err := s.UserEmailRepository.GetByAddress(&req.Email)
	if err != nil {
		return errAuthInternalServerError
	}
	if userEmail == nil {
		return nil
	}

	user, err := s.UserRepository.GetByID(&userEmail.User)
	if err != nil {
		return errAuthInternalServerError
	}
	if user == nil {
		return nil
	}

	subject := "Mais um email de teste"
	body := "<div>Testes apenas</div>"

	err = s.MailerPort.NewMessage().
		Subject(&subject).
		Body(&body).
		To(&[]string{userEmail.Address, "schilin.lucas@gmail.com"}).
		Send()
	if err != nil {
		fmt.Printf("Erro ao enviar email: %v\n", err)
		return errAuthInternalServerError
	}

	//TODO: gerar código
	//TODO: fazer envio

	return nil
}
