package dto

type AuthSignupRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type AuthSignupUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type AuthSignupResponse struct {
	User         AuthSignupUserResponse `json:"user"`
	AccessToken  string                 `json:"access_token"`
	RefreshToken string                 `json:"refresh_token"`
}
