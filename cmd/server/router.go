package main

import (
	"github.com/gorilla/mux"
	"github.com/lucasschilin/schily-users-api/internal/health"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", health.RootHandler).Methods("GET")

	return router
}
