package domain

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginRequest(email, password string) *LoginRequest {
	return &LoginRequest{Email: email, Password: password}
}
