package validator

import (
	"backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type UserValidator struct{}

func NewUserValidator() IUserValidator {
	return &UserValidator{}
}

func (uv *UserValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("Email is required"),
			validation.Length(1, 30).Error("limited max 30 characters"),
			is.Email.Error("Email is not valid"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("Password is required"),
			validation.Length(6, 30).Error("limited min 6 max 30 characters"),
		),
	)
}
