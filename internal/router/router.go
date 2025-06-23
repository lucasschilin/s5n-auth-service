package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasschilin/s5n-auth-service/internal/cache"
	"github.com/lucasschilin/s5n-auth-service/internal/handler/authhandler"
	"github.com/lucasschilin/s5n-auth-service/internal/handler/roothandler"
	"github.com/lucasschilin/s5n-auth-service/internal/middleware"
	"github.com/lucasschilin/s5n-auth-service/internal/service/authservice/jwt"
	"github.com/lucasschilin/s5n-auth-service/pkg/logger"
)

func Setup(
	l logger.Logger,
	authHand authhandler.Handler,
	rootHand roothandler.Handler,
	tokenManager jwt.TokenManager,
	cache cache.Cache,
) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.RateLimit(l, cache))
	r.Use(middleware.JSONContentType())

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
		middleware.CheckAuthentication(tokenManager)(
			(http.HandlerFunc(authHand.Validate)),
		)).Methods(http.MethodGet)

	return r
}
