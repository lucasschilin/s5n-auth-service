package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/app"
	"github.com/lucasschilin/s5n-auth-service/internal/config"
)

// @title        S5N Auth Service API
// @version      1.0
// @description  API responsável por autenticação e gerenciamento de usuários no S5N Auth Service
// @contact.name @lucasschilin @ Insta | https://github.com/lucasschilin
// @BasePath     /
func main() {
	config := config.Load()

	r := app.InitializeApp(config)

	fmt.Println("🚀 API INICIADA! ✨")
	fmt.Printf(
		"Acessível em http://%s:%s/\n\n", config.API.Host, config.API.Port,
	)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.API.Port), r))
}
