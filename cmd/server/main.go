package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucasschilin/schily-users-api/internal/config"
	"github.com/lucasschilin/schily-users-api/internal/database"
	"github.com/lucasschilin/schily-users-api/internal/router"
)

func main() {
	config := config.Load()

	r := router.New()

	database.ConnectDBUsers(config.DBUsers)
	database.ConnectDBAuth(config.DBAuth)

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
