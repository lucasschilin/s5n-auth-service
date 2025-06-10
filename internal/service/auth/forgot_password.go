package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/util"
	"github.com/lucasschilin/s5n-auth-service/internal/validator"
)

func (s *authService) ForgotPassword(req *dto.AuthForgotPasswordRequest) (
	*dto.DefaultMessageResponse, *dto.DefaultError,
) {
	messageResponse := dto.DefaultMessageResponse{
		Message: "Email sent.",
	}

	if val, detail := validator.IsValidAuthForgotPasswordRequest(req); !val {
		return nil, errorResponse(http.StatusUnprocessableEntity, detail)
	}

	if !validator.IsValidEmailAddress(req.Email) {
		return nil, errorResponse(
			http.StatusUnprocessableEntity, "Email must be a valid address",
		)
	}

	allowedRedirectHosts := []string{"s5n.com.br"}

	redirectUrlWithoutHttp := strings.ReplaceAll(strings.ToLower(req.RedirectUrl), "http://", "")
	redirectUrlWithoutHttp = strings.ReplaceAll(redirectUrlWithoutHttp, "https://", "")
	redirectUrlHost := strings.Split(redirectUrlWithoutHttp, "/")[0]

	finded, _ := util.InStringSlice(allowedRedirectHosts, redirectUrlHost)
	if !finded {
		return nil, errorResponse(
			http.StatusUnprocessableEntity,
			"Redirect URL must be a valid and allowed URL",
		)
	}

	userEmail, err := s.UserEmailRepository.GetByAddress(&req.Email)
	if err != nil {
		return nil, errAuthInternalServerError
	}
	if userEmail == nil {
		return &messageResponse, nil
	}

	user, err := s.UserRepository.GetByID(&userEmail.User)
	if err != nil {
		return nil, errAuthInternalServerError
	}
	if user == nil {
		return &messageResponse, nil
	}

	exp := time.Now().Add(5 * time.Minute).Unix()
	token, err := generateToken(s.JWTPort, "reset_password", int(exp), user.ID)
	if err != nil {
		return nil, errAuthInternalServerError
	}

	link := req.RedirectUrl + "?t=" + token

	subject := "🔐 Recupere o acesso à sua conta – redefina sua senha"
	body := fmt.Sprintf(`<p>Olá %s, como vai?</p>
			<div>
				<p>Para redefinir sua senha e recuperar sua conta, copie e cole este link no seu navegador:</p>
				<p>%s</p>
			</div>`, user.Username, link)

	err = s.MailerPort.NewMessage().
		Subject(&subject).
		Body(&body).
		To(&[]string{userEmail.Address}).
		Send()
	if err != nil {
		fmt.Printf("Erro ao enviar email: %v\n", err)
		return nil, errAuthInternalServerError
	}

	return &messageResponse, nil
}
