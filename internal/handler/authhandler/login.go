package authhandler

import (
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
)

// Login godoc
// @Summary      Login
// @Description  Realiza login e retorna tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body	dto.AuthLoginRequest  true  "Credenciais do usuário"
// @Success      200  {object}  dto.AuthLoginResponse
// @Failure      400  {object}  dto.DefaultDetailResponse
// @Failure      401  {object}  dto.DefaultDetailResponse
// @Failure      422  {object}  dto.DefaultDetailResponse
// @Failure      500  {object}  dto.DefaultDetailResponse
// @Router       /auth/login [post]
func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var req *dto.AuthLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.l.Error(err, "")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.DefaultDetailResponse{
			Detail: "The server cannot process your request.",
		})
		return
	}

	res, err := h.AuthService.Login(h.l, req)
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
