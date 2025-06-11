package authservice

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aidarkhanov/nanoid"
	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Signup(req *dto.AuthSignupRequest) (
	*dto.AuthLoginResponse, *dto.DefaultError,
) {
	if val, detail := validator.IsValidAuthSignupRequest(req); !val {
		return nil, errorResponse(http.StatusUnprocessableEntity, detail)
	}

	if len(req.Password) < MinPasswordLength {
		return nil, errorResponse(
			http.StatusUnprocessableEntity,
			fmt.Sprintf(
				"Password must have at least %v characters.",
				MinPasswordLength,
			),
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
