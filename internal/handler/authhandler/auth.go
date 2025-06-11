package authhandler

import (
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/service/authservice"
)

type Handler interface {
	Signup(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	ForgotPassword(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
	Validate(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	AuthService authservice.Service
}

func NewHandler(authServ authservice.Service) Handler {
	return &handler{
		AuthService: authServ,
	}
}
