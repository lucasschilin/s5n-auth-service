package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}
	apiPort := os.Getenv("API_PORT")
	apiHost := os.Getenv("API_HOST")

	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(map[string]string{
			"detail": "Olá Mundo, esta é a SCHILY USERS API!",
		})
	}).Methods("GET")

	// Cores ANSI para o terminal
	green := "\033[32m"
	yellow := "\033[33m"
	blue := "\033[34m"
	reset := "\033[0m"

	fmt.Printf("%s🚀 API INICIADA! 🚀%s\n", green, reset)
	fmt.Printf("%sAcessível em http://%s:%s%s/\n", yellow, apiHost, apiPort, reset)
	fmt.Printf("%sAPI rodando... ✨ 🌐%s\n", blue, reset)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", apiPort), router))
}
