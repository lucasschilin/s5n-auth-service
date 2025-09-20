package authhandler

import (
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
)

// Refresh godoc
// @Summary      Renovar access token
// @Description  Gera um novo access token a partir de um refresh token válido
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AuthRefreshRequest  true  "Refresh token"
// @Success      200  {object}  dto.AuthRefreshResponse
// @Failure      400  {object}  dto.DefaultDetailResponse "Requisição malformada"
// @Failure      401  {object}  dto.DefaultDetailResponse "Refresh token inválido ou expirado"
// @Failure      422  {object}  dto.DefaultDetailResponse "Payload inválido"
// @Failure      500  {object}  dto.DefaultDetailResponse "Erro interno do servidor"
// @Router       /auth/refresh [post]
func (h *handler) Refresh(w http.ResponseWriter, r *http.Request) {
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

// TODO: Log refresh tokens
