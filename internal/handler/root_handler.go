package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
)

type RootHandler interface {
	Root(w http.ResponseWriter, r *http.Request)
}

type rootHandler struct{}

func NewRootHandler() RootHandler {
	return &rootHandler{}
}

func (h *rootHandler) Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.DefaultMessageResponse{
		Message: "S5N Auth Service API healthed and online 🟢",
	})
}
