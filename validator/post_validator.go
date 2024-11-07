package validator

import (
	"backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IPostValidator interface {
	PostValidate(post model.Post) error
}

type PostValidator struct{}

func NewPostValidator() IPostValidator {
	return &PostValidator{}
}

func (pv *PostValidator) PostValidate(post model.Post) error {
	return validation.ValidateStruct(&post,
		validation.Field(
			&post.Content,
			validation.Required.Error("Content is required"),
			validation.Length(1, 15).Error("limited max 15 characters"),
		),
	)
}
