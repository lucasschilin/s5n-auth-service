package router

import (
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	RegisterRootRoutes(r)

	return r
}
