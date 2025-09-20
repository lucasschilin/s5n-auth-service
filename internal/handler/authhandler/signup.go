package authhandler

import (
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
)

// Signup godoc
// @Summary      Criar usuário
// @Description  Cria uma nova conta de usuário e retorna tokens de autenticação
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AuthSignupRequest  true  "Dados do usuário"
// @Success      200  {object}  dto.AuthLoginResponse
// @Failure      400  {object}  dto.DefaultDetailResponse "JSON inválido"
// @Failure      422  {object}  dto.DefaultDetailResponse "Dados inválidos (email ou senha)"
// @Failure      409  {object}  dto.DefaultDetailResponse "Email já em uso"
// @Failure      500  {object}  dto.DefaultDetailResponse "Erro interno do servidor"
// @Router       /auth/signup [post]
func (h *handler) Signup(w http.ResponseWriter, r *http.Request) {
	var req *dto.AuthSignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.l.Error(err, "")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.DefaultDetailResponse{
			Detail: "The server cannot process your request.",
		})
		return
	}

	res, err := h.AuthService.Signup(h.l, req)
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
