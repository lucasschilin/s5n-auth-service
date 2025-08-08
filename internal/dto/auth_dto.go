package dto

type AuthSignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	User struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthRefreshResponse struct {
	AccessToken string `json:"access_token"`
}

type AuthForgotPasswordRequest struct {
	Email       string `json:"email"`
	RedirectUrl string `json:"redirect_url"`
}

type AuthResetPasswordRequest struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type AuthValidateResponse struct {
	User struct {
		ID string `json:"id"`
	} `json:"user"`
}
