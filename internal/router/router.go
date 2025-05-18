package router

import (
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	registerRoutes(r)

	return r
}

func registerRoutes(r *mux.Router) {
	registerRootRoutes(r)
	registerAuthRoutes(r)
}
