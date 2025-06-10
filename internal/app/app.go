package app

import (
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/adapter"
	"github.com/lucasschilin/s5n-auth-service/internal/config"
	"github.com/lucasschilin/s5n-auth-service/internal/database"
	"github.com/lucasschilin/s5n-auth-service/internal/handler"
	"github.com/lucasschilin/s5n-auth-service/internal/repository"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/password"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/user"
	"github.com/lucasschilin/s5n-auth-service/internal/router"
	"github.com/lucasschilin/s5n-auth-service/internal/service/auth"
)

func InitializeApp(config *config.Config) http.Handler {
	usersDB := database.ConnectDBUsers(config.DBUsers)
	authDB := database.ConnectDBAuth(config.DBAuth)

	userRepo := user.NewRepository(usersDB)
	userEmailRepo := repository.NewUserEmailRepository(usersDB)
	passwordRepo := password.NewRepository(authDB)

	jwtAdapter := adapter.NewJWT(config.JWT.SecretKey)
	mailerAdapter := adapter.NewSMTPMailer(
		&config.SMTP.Host, &config.SMTP.Port, &config.SMTP.Username,
		&config.SMTP.Password, &config.SMTP.From,
	)

	authServ := auth.NewService(
		usersDB, authDB, userRepo, userEmailRepo,
		passwordRepo, jwtAdapter, mailerAdapter,
	)

	authHand := handler.NewAuthHandler(authServ)
	rootHand := handler.NewRootHandler()

	r := router.Setup(authHand, rootHand, jwtAdapter)

	return r
}
