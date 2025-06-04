package app

import (
	"net/http"

	"github.com/lucasschilin/schily-users-api/internal/adapter"
	"github.com/lucasschilin/schily-users-api/internal/config"
	"github.com/lucasschilin/schily-users-api/internal/database"
	"github.com/lucasschilin/schily-users-api/internal/handler"
	"github.com/lucasschilin/schily-users-api/internal/repository"
	"github.com/lucasschilin/schily-users-api/internal/router"
	"github.com/lucasschilin/schily-users-api/internal/service"
)

func InitializeApp(config *config.Config) http.Handler {
	usersDB := database.ConnectDBUsers(config.DBUsers)
	authDB := database.ConnectDBAuth(config.DBAuth)

	userRepo := repository.NewUserRepository(usersDB)
	userEmailRepo := repository.NewUserEmailRepository(usersDB)
	passwordRepo := repository.NewPasswordRepository(authDB)

	jwtAdapter := adapter.NewJWT(config.JWT.SecretKey)
	mailerAdapter := adapter.NewSMTPMailer(
		&config.SMTP.Host, &config.SMTP.Port, &config.SMTP.Username,
		&config.SMTP.Password, &config.SMTP.From,
	)

	authServ := service.NewAuthService(
		usersDB, authDB, userRepo, userEmailRepo,
		passwordRepo, jwtAdapter, mailerAdapter,
	)

	authHand := handler.NewAuthHandler(authServ)
	rootHand := handler.NewRootHandler()

	r := router.Setup(authHand, rootHand)

	return r
}
