package auth

import (
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/middleware"
)

func (h *handler) Validate(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(string)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.AuthValidateResponse{
		User: struct {
			ID string `json:"id"`
		}{
			ID: userID,
		},
	})
}
