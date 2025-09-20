package authhandler

import (
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/middleware"
)

// Validate godoc
// @Summary      Validar token de autenticação
// @Description  Verifica se o access token informado é válido e retorna os dados do usuário
// @Tags         auth
// @Produce      json
// @Param        Authorization  header	string  false  "Access token no formato: Bearer {token}"
// @Success      200  {object}  dto.AuthValidateResponse
// @Failure      401  {object}  dto.DefaultDetailResponse "Token inválido ou ausente"
// @Failure      500  {object}  dto.DefaultDetailResponse "Erro interno do servidor"
// @Router       /auth/validate [get]
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

// TODO: Log validate password
