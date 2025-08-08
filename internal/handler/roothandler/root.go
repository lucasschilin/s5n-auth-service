package roothandler

import (
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/pkg/logger"
)

type Handler interface {
	Root(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	l logger.Logger
}

func NewHandler(l logger.Logger) Handler {
	return &handler{
		l: l,
	}
}

func (h *handler) Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.DefaultMessageResponse{
		Message: "S5N Auth Service API healthed and online 🟢",
	})
}
