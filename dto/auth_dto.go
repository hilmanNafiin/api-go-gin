package dto

type RegisterRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Gender string `json:"gender"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}