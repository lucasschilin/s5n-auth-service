package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/app"
	"github.com/lucasschilin/s5n-auth-service/internal/config"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server TESTINHUUU server.

func main() {
	config := config.Load()

	r := app.InitializeApp(config)

	fmt.Println("🚀 API INICIADA! ✨")
	fmt.Printf(
		"Acessível em http://%s:%s/\n\n", config.API.Host, config.API.Port,
	)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.API.Port), r))
}
