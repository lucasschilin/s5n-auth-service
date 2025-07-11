package authhandler

import (
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/service/authservice"
	"github.com/lucasschilin/s5n-auth-service/pkg/logger"
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
	l           logger.Logger
	AuthService authservice.Service
}

func NewHandler(l logger.Logger, authServ authservice.Service) Handler {
	return &handler{
		l:           l,
		AuthService: authServ,
	}
}
