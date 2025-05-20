package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/schily-users-api/internal/dto"
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
	var req *dto.AuthSignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.DefaultDetailResponse{
			Detail: err.Error(),
		})
		return
	}

	res, err := h.AuthService.Signup(req)
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(dto.DefaultDetailResponse{
			Detail: err.Detail,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
