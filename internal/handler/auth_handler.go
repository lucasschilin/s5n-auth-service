package handler

import (
	"fmt"
	"net/http"

	"github.com/lucasschilin/schily-users-api/internal/service"
)

type AuthHandler interface {
	Signup(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(authServ service.AuthService) AuthHandler {
	return &authHandler{
		AuthService: authServ,
	}
}

func (h *authHandler) Signup(w http.ResponseWriter, r *http.Request) {
	//

	fmt.Println("teste")
}
