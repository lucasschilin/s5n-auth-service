package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/aidarkhanov/nanoid"
	"github.com/lucasschilin/schily-users-api/internal/dto"
	"github.com/lucasschilin/schily-users-api/internal/repository"
	"github.com/lucasschilin/schily-users-api/internal/validator"
	"golang.org/x/crypto/bcrypt"
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
			Detail: "Email cannot be validated.",
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
	fmt.Println(newUser)

	verifyToken, err := nanoid.Generate(nanoid.DefaultAlphabet, 50)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred.",
		}
	}

	newUserEmail, err := s.UserEmailRepository.CreateWithTX(
		usersTX, &newUser.ID, &req.Email, &verifyToken,
	)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred.",
		}
	}
	fmt.Println(newUserEmail)

	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred.",
		}
	}

	newPassword, err := s.PasswordRepository.CreateWithTX(
		authTX, &newUser.ID, string(bcryptedPassword),
	)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: "An error occurred. " + err.Error(),
		}
	}
	fmt.Println(newPassword)
	// TODO: generate JWT

	// usersTX.Commit()
	// authTX.Commit()

	return &dto.AuthSignupResponse{
		User: dto.AuthSignupUserResponse{
			ID:       newUser.ID,
			Username: newUser.Username,
		},
		AccessToken:  "accesstoketeste",
		RefreshToken: "refreshoketeste",
	}, nil

}

// TODO: refactor return error (DRY)
