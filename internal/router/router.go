package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasschilin/schily-users-api/internal/handler"
)

func Setup(authHand handler.AuthHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/auth/signup", authHand.Signup).Methods(http.MethodPost)

	return r
}
