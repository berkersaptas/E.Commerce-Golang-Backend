package models

type UserForgetPasswordModel struct {
	Email string `json:"email,omitempty" validate:"required"`
}
