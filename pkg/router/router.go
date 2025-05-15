package router

import (
	"github.com/gorilla/mux"
	"github.com/lucasschilin/schily-users-api/internal/root"
)

func New() *mux.Router {
	r := mux.NewRouter()

	root.RegisterRoutes(r)

	return r
}
