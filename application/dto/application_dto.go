package dto

type UserInput struct {
	ID       string
	Name     string
	Email    string
	Password string
	Role     string
	Status   bool
}

type AuthenticateOutput struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
