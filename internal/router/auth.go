package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasschilin/schily-users-api/internal/handler"
)

func registerAuthRoutes(r *mux.Router) {
	apiRouter := r.PathPrefix("/auth").Subrouter()

	apiRouter.HandleFunc("/signup", handler.AuthSignup).Methods(http.MethodPost)
}
