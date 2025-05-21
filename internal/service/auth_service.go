package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/aidarkhanov/nanoid"
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

	userEmail, err := s.UserEmailRepository.GetByAddress(req.Email)
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

	newID := nanoid.New()
	username := "u_" + newID

	user, err := s.UserRepository.CreateWithTX(usersTX, newID, username)
	if err != nil {
		return nil, &dto.DefaultError{
			Code:   http.StatusInternalServerError,
			Detail: err.Error(),
			// Detail: "An error occurred. ",
		}
	}
	fmt.Println(user)

	// TODO: create user_email
	// TODO: generate uuid
	// TODO: crypt password
	// TODO: create password

	return &dto.AuthSignupResponse{
		User: dto.AuthSignupUserResponse{
			ID:       "tsaufsb",
			Username: "lucaslash",
		},
		AccessToken:  "accesstoketeste",
		RefreshToken: "refreshoketeste",
	}, nil

}

// TODO: refactor return error (DRY)
