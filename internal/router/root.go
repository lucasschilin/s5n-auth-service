package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasschilin/schily-users-api/internal/handler"
)

func registerRootRoutes(r *mux.Router) {
	r.HandleFunc("/", handler.RootHandler).Methods(http.MethodGet)
}
