package app

import (
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/adapter"
	"github.com/lucasschilin/s5n-auth-service/internal/config"
	"github.com/lucasschilin/s5n-auth-service/internal/database"
	authhandler "github.com/lucasschilin/s5n-auth-service/internal/handler/auth"
	"github.com/lucasschilin/s5n-auth-service/internal/handler/root"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/password"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/user"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/useremail"
	"github.com/lucasschilin/s5n-auth-service/internal/router"
	authservice "github.com/lucasschilin/s5n-auth-service/internal/service/auth"
)

func InitializeApp(config *config.Config) http.Handler {
	usersDB := database.ConnectDBUsers(config.DBUsers)
	authDB := database.ConnectDBAuth(config.DBAuth)

	userRepo := user.NewRepository(usersDB)
	userEmailRepo := useremail.NewRepository(usersDB)
	passwordRepo := password.NewRepository(authDB)

	jwtAdapter := adapter.NewJWT(config.JWT.SecretKey)
	mailerAdapter := adapter.NewSMTPMailer(
		&config.SMTP.Host, &config.SMTP.Port, &config.SMTP.Username,
		&config.SMTP.Password, &config.SMTP.From,
	)

	authServ := authservice.NewService(
		usersDB, authDB, userRepo, userEmailRepo,
		passwordRepo, jwtAdapter, mailerAdapter,
	)

	authHand := authhandler.NewHandler(authServ)
	rootHand := root.NewHandler()

	r := router.Setup(authHand, rootHand, jwtAdapter)

	return r
}
