package dto

type UserInput struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
}

type AuthenticateOutput struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
