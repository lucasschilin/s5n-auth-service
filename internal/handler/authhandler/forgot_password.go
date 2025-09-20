package authhandler

import (
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
)

// ForgotPassword godoc
// @Summary      Esqueci minha senha
// @Description  Inicia o fluxo de recuperação de senha enviando um e-mail com link de redefinição
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AuthForgotPasswordRequest  true  "Dados do usuário e redirect URL"
// @Success      200  {object}  dto.DefaultMessageResponse "Mensagem de sucesso"
// @Failure      400  {object}  dto.DefaultDetailResponse "Requisição malformada"
// @Failure      422  {object}  dto.DefaultDetailResponse "Dados inválidos ou redirect_url não permitido"
// @Failure      500  {object}  dto.DefaultDetailResponse "Erro interno do servidor"
// @Router       /auth/forgot-password [post]
func (h *handler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
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

// TODO: Log forgot password
