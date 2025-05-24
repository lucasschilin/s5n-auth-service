package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aidarkhanov/nanoid"
	"golang.org/x/crypto/bcrypt"

	"github.com/lucasschilin/schily-users-api/internal/adapter"
	"github.com/lucasschilin/schily-users-api/internal/config"
	"github.com/lucasschilin/schily-users-api/internal/dto"
	"github.com/lucasschilin/schily-users-api/internal/repository"
	"github.com/lucasschilin/schily-users-api/internal/validator"
)

type AuthService interface {
	Signup(req *dto.AuthSignupRequest) (
		*dto.AuthSignupResponse, *dto.DefaultError,
	)
}

type authService struct {
	UsersDB             *sql.DB
	AuthDB              *sql.DB
	UserRepository      repository.UserRepository
	UserEmailRepository repository.UserEmailRepository
	PasswordRepository  repository.PasswordRepository
}

func NewAuthService(
	usersDB *sql.DB,
	authDB *sql.DB,
	userRepo repository.UserRepository,
	userEmailRepo repository.UserEmailRepository,
	passwordRepo repository.PasswordRepository,
) AuthService {
	return &authService{
		UsersDB:             usersDB,
		AuthDB:              authDB,
		UserRepository:      userRepo,
		UserEmailRepository: userEmailRepo,
		PasswordRepository:  passwordRepo,
	}
}

func (s *authService) Signup(req *dto.AuthSignupRequest) (
	*dto.AuthSignupResponse, *dto.DefaultError,
) {
	if val, detail := validator.IsValidAuthSignupRequest(req); !val {
		return nil, errorResponse(
			http.StatusUnprocessableEntity,
			detail,
		)
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
			http.StatusBadRequest, "Email must be a valid address",
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
		return nil, errAuthSignupInternalServerError
	}

	authTX, err := s.AuthDB.Begin()
	if err != nil {
		return nil, errorResponse(
			http.StatusInternalServerError, "An error occurred.",
		)
	}

	defer usersTX.Rollback()
	defer authTX.Rollback()

	emailUsername := strings.Split(req.Email, "@")[0]

	maxUsernameLength := 13
	if len(emailUsername) < maxUsernameLength {
		maxUsernameLength = len(emailUsername)
	}
	username := strings.Replace(
		strings.ToLower(emailUsername[:maxUsernameLength]), ".", "", -1,
	)

	user, err := s.UserRepository.GetByUsername(&username)
	if err != nil {
		return nil, errAuthSignupInternalServerError
	}
	if user != nil {
		sufix, err := nanoid.Generate(nanoid.DefaultAlphabet, 5)
		if err != nil {
			return nil, errAuthSignupInternalServerError
		}
		username = strings.ToLower(fmt.Sprintf("%s_%s", username, sufix))
	}

	newUser, err := s.UserRepository.CreateWithTX(usersTX, &username)
	if err != nil {
		return nil, errAuthSignupInternalServerError
	}

	verifyToken, err := nanoid.Generate(nanoid.DefaultAlphabet, 50)
	if err != nil {
		return nil, errAuthSignupInternalServerError
	}

	_, err = s.UserEmailRepository.CreateWithTX(
		usersTX, &newUser.ID, &req.Email, &verifyToken,
	)
	if err != nil {
		return nil, errAuthSignupInternalServerError
	}

	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return nil, errAuthSignupInternalServerError
	}

	_, err = s.PasswordRepository.CreateWithTX(
		authTX, &newUser.ID, string(bcryptedPassword),
	)
	if err != nil {
		return nil, errAuthSignupInternalServerError
	}

	config := config.Load()
	jwtSecretKey := config.JWT.SecretKey

	jwtAdapter := adapter.NewJWT(jwtSecretKey)

	accessTokenExpiration := time.Now().Add(30 * time.Minute).Unix()
	mapClaims := map[string]interface{}{
		"iat":  time.Now().Unix(),
		"exp":  accessTokenExpiration,
		"sub":  newUser.ID,
		"type": "access_token",
	}
	accessToken, err := jwtAdapter.GenerateToken(mapClaims)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errAuthSignupInternalServerError
	}

	refreshTokenExpiration := time.Now().Add(24 * time.Hour).Unix()
	mapClaims = map[string]interface{}{
		"iat":  time.Now().Unix(),
		"exp":  refreshTokenExpiration,
		"sub":  newUser.ID,
		"type": "refresh_token",
	}
	refreshToken, err := jwtAdapter.GenerateToken(mapClaims)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errAuthSignupInternalServerError
	}

	usersTX.Commit()
	authTX.Commit()

	return &dto.AuthSignupResponse{
		User: dto.AuthSignupUserResponse{
			ID:       newUser.ID,
			Username: newUser.Username,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
