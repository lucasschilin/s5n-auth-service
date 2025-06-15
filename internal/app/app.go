package app

import (
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/adapter"
	"github.com/lucasschilin/s5n-auth-service/internal/cache"
	"github.com/lucasschilin/s5n-auth-service/internal/config"
	"github.com/lucasschilin/s5n-auth-service/internal/database"
	"github.com/lucasschilin/s5n-auth-service/internal/handler/authhandler"
	"github.com/lucasschilin/s5n-auth-service/internal/handler/roothandler"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/password"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/user"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/useremail"
	"github.com/lucasschilin/s5n-auth-service/internal/router"
	"github.com/lucasschilin/s5n-auth-service/internal/service/authservice"
)

func InitializeApp(config *config.Config) http.Handler {
	usersDB := database.ConnectDBUsers(config.DBUsers)
	authDB := database.ConnectDBAuth(config.DBAuth)

	cache := cache.NewRedisClient(
		config.Redis.Host, config.Redis.Port, config.Redis.Password, 0,
	)

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
	rootHand := roothandler.NewHandler()

	r := router.Setup(authHand, rootHand, jwtAdapter, cache)

	return r
}
