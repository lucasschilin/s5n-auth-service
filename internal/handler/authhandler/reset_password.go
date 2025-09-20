package authhandler

import (
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
)

// ResetPassword godoc
// @Summary      Redefinir senha
// @Description  Redefine a senha do usuário usando o token recebido por e-mail
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body	dto.AuthResetPasswordRequest  true  "Token de redefinição e nova senha"
// @Success      200  {object}  dto.DefaultMessageResponse "Senha redefinida com sucesso"
// @Failure      400  {object}  dto.DefaultDetailResponse "Requisição malformada"
// @Failure      422  {object}  dto.DefaultDetailResponse "Dados inválidos ou regras de senha violadas"
// @Failure      401  {object}  dto.DefaultDetailResponse "Token inválido ou expirado"
// @Failure      500  {object}  dto.DefaultDetailResponse "Erro interno do servidor"
// @Router       /auth/reset-password [post]
func (h *handler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req *dto.AuthResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.DefaultDetailResponse{
			Detail: "The server cannot process your request.",
		})
		return
	}

	res, err := h.AuthService.ResetPassword(req)
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

// TODO: Log reset password
