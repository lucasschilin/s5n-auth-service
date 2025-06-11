package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasschilin/s5n-auth-service/internal/handler/authhandler"
	"github.com/lucasschilin/s5n-auth-service/internal/handler/roothandler"
	"github.com/lucasschilin/s5n-auth-service/internal/middleware"
	"github.com/lucasschilin/s5n-auth-service/internal/port"
)

func Setup(
	authHand authhandler.Handler,
	rootHand roothandler.Handler,
	jwtPort port.JWT,
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
	auth.Handle(
		"/validate",
		middleware.CheckAuthentication(
			http.HandlerFunc(authHand.Validate),
			jwtPort,
		)).Methods(http.MethodGet)

	return r
}
