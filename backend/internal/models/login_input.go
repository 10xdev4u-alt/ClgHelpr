package models

// LoginUserInput defines the expected input for user login.
type LoginUserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
