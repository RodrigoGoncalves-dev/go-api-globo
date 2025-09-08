package domain

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Token   string `json:"token"`
	Expires string `json:"expires_at"`
}

type IAuthRepository interface {
	DoLogin(input LoginInput) (*LoginOutput, error)
}
