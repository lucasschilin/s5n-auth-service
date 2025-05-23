package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aidarkhanov/nanoid"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

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
	if _, err := validator.IsValidAuthSignupRequest(req); err != nil {
		return nil, err
	}

	const MinPasswordLength = 8
	if len(req.Password) < MinPasswordLength {
		return nil, &dto.DefaultError{
			Code: http.StatusUnprocessableEntity,
			Detail: fmt.Sprintf(
				"Password must have at least %v characters.",
				MinPasswordLength,
			),
		}
	}

	if req.Password != req.ConfirmPassword {
		return nil, &dto.DefaultError{
			Code:   http.StatusUnprocessableEntity,
			Detail: "Password and confirmation password must match.",
		}
	}

	if !validator.IsValidEmailAddress(req.Email) {
		return nil, &dto.DefaultError{
			Code:   http.StatusBadRequest,
			Detail: "Email must be a valid address",
		}
	}

	userEmail, err := s.UserEmailRepository.GetByAddress(&req.Email)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "Email cannot be validated. " + err.Error(),
		}
	}
	if userEmail != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusConflict,
			Detail: "Email already in use.",
		}
	}

	usersTX, err := s.UsersDB.Begin()
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred.",
		}
	}

	authTX, err := s.AuthDB.Begin()
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred.",
		}
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
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred." + err.Error(),
		}
	}
	if user != nil {
		sufix, err := nanoid.Generate(nanoid.DefaultAlphabet, 5)
		if err != nil {
			return nil, &dto.DefaultError{
				Code:   http.StatusInternalServerError,
				Detail: "An error occurred.",
			}
		}
		username = strings.ToLower(fmt.Sprintf("%s_%s", username, sufix))
	}

	newUser, err := s.UserRepository.CreateWithTX(usersTX, &username)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred.",
		}
	}

	verifyToken, err := nanoid.Generate(nanoid.DefaultAlphabet, 50)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred.",
		}
	}

	_, err = s.UserEmailRepository.CreateWithTX(
		usersTX, &newUser.ID, &req.Email, &verifyToken,
	)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred.",
		}
	}

	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred.",
		}
	}

	_, err = s.PasswordRepository.CreateWithTX(
		authTX, &newUser.ID, string(bcryptedPassword),
	)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred.",
		}
	}

	config := config.Load()

	secretKey := []byte(config.JWT.SecretKey)
	accessTokenExpiration := time.Now().Add(30 * time.Minute)
	refreshTokenExpiration := time.Now().Add(24 * time.Hour)

	mapClaims := jwt.MapClaims{
		"iat":  time.Now().Unix(),
		"exp":  accessTokenExpiration,
		"sub":  newUser.ID,
		"type": "access_token",
	}
	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	accessToken, err := accessTokenObj.SignedString(secretKey)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred. " + err.Error(),
		}
	}

	mapClaims = jwt.MapClaims{
		"iat":  time.Now().Unix(),
		"exp":  refreshTokenExpiration,
		"sub":  newUser.ID,
		"type": "refresh_token",
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	refreshToken, err := refreshTokenObj.SignedString(secretKey)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred. " + err.Error(),
		}
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

// TODO: refactor return errors (DRY)
// TODO: migrate bcrypt, JWT to adapters
