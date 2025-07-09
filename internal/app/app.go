package app

import (
	"net/http"
	"strconv"

	"github.com/lucasschilin/s5n-auth-service/internal/cache"
	"github.com/lucasschilin/s5n-auth-service/internal/config"
	"github.com/lucasschilin/s5n-auth-service/internal/database"
	"github.com/lucasschilin/s5n-auth-service/internal/handler/authhandler"
	"github.com/lucasschilin/s5n-auth-service/internal/handler/roothandler"
	"github.com/lucasschilin/s5n-auth-service/internal/integrations/mailer"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/password"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/user"
	"github.com/lucasschilin/s5n-auth-service/internal/repository/useremail"
	"github.com/lucasschilin/s5n-auth-service/internal/router"
	"github.com/lucasschilin/s5n-auth-service/internal/service/authservice"
	"github.com/lucasschilin/s5n-auth-service/internal/service/authservice/jwt"
	"github.com/lucasschilin/s5n-auth-service/pkg/logger"
)

func InitializeApp(config *config.Config) http.Handler {
	logLevel, _ := strconv.ParseInt(config.Log.Level, 10, 8)
	l := logger.New(int(logLevel))

	usersDB := database.ConnectDBUsers(l, config.DBUsers)
	authDB := database.ConnectDBAuth(l, config.DBAuth)

	cache := cache.NewRedisClient(
		config.Redis.Host, config.Redis.Port, config.Redis.Password, 0,
	)

	userRepo := user.NewRepository(usersDB)
	userEmailRepo := useremail.NewRepository(usersDB)
	passwordRepo := password.NewRepository(authDB)

	jwtManager := jwt.NewJWT(config.JWT.SecretKey)
	mailerAdapter := mailer.NewSMTPMailer(
		&config.SMTP.Host, &config.SMTP.Port, &config.SMTP.Username,
		&config.SMTP.Password, &config.SMTP.From,
	)

	authServ := authservice.NewService(
		usersDB, authDB, userRepo, userEmailRepo,
		passwordRepo, jwtManager, mailerAdapter,
	)

	authHand := authhandler.NewHandler(authServ)
	rootHand := roothandler.NewHandler()

	r := router.Setup(l, authHand, rootHand, jwtManager, cache)

	return r
}
