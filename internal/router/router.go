package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasschilin/s5n-auth-service/internal/handler"
	"github.com/lucasschilin/s5n-auth-service/internal/middleware"
)

func Setup(
	authHand handler.AuthHandler,
	rootHand handler.RootHandler,
) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.JSONContentType)

	root := r.PathPrefix("").Subrouter()
	root.HandleFunc("/", rootHand.Root).Methods(http.MethodGet)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", authHand.Signup).Methods(http.MethodPost)
	auth.HandleFunc("/login", authHand.Login).Methods(http.MethodPost)
	auth.HandleFunc("/refresh", authHand.Refresh).Methods(http.MethodPost)
	auth.HandleFunc("/forgot-password", authHand.ForgotPassword).Methods(http.MethodPost)
	auth.HandleFunc("/reset-password", authHand.ResetPassword).Methods(http.MethodPost)

	requireLogin := auth.PathPrefix("").Subrouter()
	requireLogin.Use(middleware.CheckAuthentication)

	requireLogin.HandleFunc("/validate", authHand.Validate).Methods(http.MethodGet)

	return r
}
