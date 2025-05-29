package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasschilin/schily-users-api/internal/handler"
)

func Setup(
	authHand handler.AuthHandler,
	rootHand handler.RootHandler,
) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHand.Root).Methods(http.MethodGet)

	r.HandleFunc("/auth/signup", authHand.Signup).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", authHand.Login).Methods(http.MethodPost)

	return r
}
