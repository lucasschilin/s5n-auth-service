package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/schily-users-api/internal/dto"
)

type RootHandler interface {
	Root(w http.ResponseWriter, r *http.Request)
}

type rootHandler struct {
	DB *sql.DB
}

func (h *rootHandler) NewRootHandler(db *sql.DB) RootHandler {
	return &rootHandler{
		DB: db,
	}
}

func (h *rootHandler) Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.DefaultMessageResponse{
		Message: "Schily Users API healthed and online 🟢",
	})
}
