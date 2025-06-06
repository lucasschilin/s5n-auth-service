package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/schily-users-api/internal/dto"
	"github.com/lucasschilin/schily-users-api/internal/service"
)

type AuthHandler interface {
	Signup(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	ForgotPassword(w http.ResponseWriter, r *http.Request)
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.DefaultDetailResponse{
			Detail: "The server cannot process your request.",
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

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req *dto.AuthLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.DefaultDetailResponse{
			Detail: "The server cannot process your request.",
		})
		return
	}

	res, err := h.AuthService.Login(req)
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

func (h *authHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req *dto.AuthRefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.DefaultDetailResponse{
			Detail: "The server cannot process your request.",
		})
		return
	}

	res, err := h.AuthService.Refresh(req)
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

func (h *authHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var req *dto.AuthForgotPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.DefaultDetailResponse{
			Detail: "The server cannot process your request.",
		})
		return
	}

	res, err := h.AuthService.ForgotPassword(req)
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
