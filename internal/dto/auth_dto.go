package dto

// AuthSignupRequest representa o payload para criar um novo usuário na aplicação
type AuthSignupRequest struct {
	// Email do usuário
	// Deve ser único e válido
	// example: lucas@schilin.com
	Email string `json:"email" example:"lucas@schilin.com"`
	// Senha do usuário (não existe regras de formatação e será hasheada antes de persistir)
	// example: 123456
	Password string `json:"password" example:"12345678"`
}

// AuthLoginRequest representa o payload de autenticação na aplicação
type AuthLoginRequest struct {
	// Email do usuário já cadastrado
	// example: lucas@schilin.com
	Email string `json:"email" example:"lucas@schilin.com"`
	// Senha em texto plano que será validada
	// example: 123456
	Password string `json:"password" example:"12345678"`
}

// AuthLoginResponse representa o retorno em caso de sucesso na autenticação
type AuthLoginResponse struct {
	// Informações básicas do usuário autenticado
	User struct {
		// Identificador único do usuário
		// example: 123e4567-e89b-12d3-a456-426614174000
		ID string `json:"id"`
		// Nome de exibição do usuário
		// example: Lucas
		Username string `json:"username"`
	} `json:"user"`
	// Token JWT válido para acessar recursos protegidos
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	AccessToken string `json:"access_token"`
	// Token de atualização (refresh) para gerar novos access tokens
	// example: def502009c3b9b...
	RefreshToken string `json:"refresh_token"`
}

// AuthRefreshRequest representa o payload para renovar o access token
type AuthRefreshRequest struct {
	// Refresh token válido emitido anteriormente
	// example: def502009c3b9b...
	RefreshToken string `json:"refresh_token" example:"def502009c3b9b..."`
}

// AuthRefreshResponse representa a resposta ao renovar o access token
type AuthRefreshResponse struct {
	// Novo token de acesso
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	AccessToken string `json:"access_token"`
}

// AuthForgotPasswordRequest representa o payload para iniciar o fluxo de recuperação de senha
type AuthForgotPasswordRequest struct {
	// Email do usuário que esqueceu a senha
	// example: lucas@schilin.com
	Email string `json:"email" example:"lucas@schilin.com"`
	// URL para onde o usuário será redirecionado após redefinir a senha (precisa estar na lista de URLs permitidas)
	// example: https://app.exemplo.com/reset
	RedirectUrl string `json:"redirect_url" example:"https://s5n.com.br/reset"`
}

// AuthResetPasswordRequest representa o payload para redefinir a senha
type AuthResetPasswordRequest struct {
	// Token temporário recebido por e-mail
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	// Nova senha que substituirá a antiga
	// example: novasenha12345678
	NewPassword string `json:"new_password" example:"novasenha12345678"`
}

// AuthValidateResponse representa a resposta de validação de um token de autenticação
type AuthValidateResponse struct {
	// Dados do usuário validados
	User struct {
		// Identificador único do usuário
		// example: 123e4567-e89b-12d3-a456-426614174000
		ID string `json:"id"`
	} `json:"user"`
}
