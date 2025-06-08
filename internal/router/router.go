package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasschilin/s5n-auth-service/internal/handler"
)

func Setup(
	authHand handler.AuthHandler,
	rootHand handler.RootHandler,
) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHand.Root).Methods(http.MethodGet)

	r.HandleFunc("/auth/signup", authHand.Signup).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", authHand.Login).Methods(http.MethodPost)
	r.HandleFunc("/auth/refresh", authHand.Refresh).Methods(http.MethodPost)
	r.HandleFunc("/auth/forgot-password", authHand.ForgotPassword).Methods(http.MethodPost)
	r.HandleFunc("/auth/reset-password", authHand.ResetPassword).Methods(http.MethodPost)

	return r
}
