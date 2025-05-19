package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucasschilin/schily-users-api/internal/config"
	"github.com/lucasschilin/schily-users-api/internal/database"
	"github.com/lucasschilin/schily-users-api/internal/handler"
	"github.com/lucasschilin/schily-users-api/internal/repository"
	"github.com/lucasschilin/schily-users-api/internal/router"
	"github.com/lucasschilin/schily-users-api/internal/service"
)

func main() {
	config := config.Load()

	usersDB := database.ConnectDBUsers(config.DBUsers)
	authDB := database.ConnectDBAuth(config.DBAuth)

	userRepo := repository.NewUserRepository(usersDB)
	userEmailRepo := repository.NewUserEmailRepository(usersDB)
	passwordRepo := repository.NewPasswordRepository(authDB)

	authServ := service.NewAuthService(userRepo, userEmailRepo, passwordRepo)

	authHand := handler.NewAuthHandler(authServ)

	r := router.Setup(authHand)

	// Cores ANSI para o terminal
	green := "\033[32m"
	yellow := "\033[33m"
	blue := "\033[34m"
	reset := "\033[0m"

	fmt.Printf("%s🚀 API INICIADA! 🚀%s\n", green, reset)
	fmt.Printf("%sAcessível em http://%s:%s%s/\n", yellow, config.API.Host, config.API.Port, reset)
	fmt.Printf("%sAPI rodando... ✨ 🌐%s\n", blue, reset)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.API.Port), r))
}
