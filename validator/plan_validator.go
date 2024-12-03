package validator

import (
	"backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IPlanValidator interface {
	PlanValidate(plan model.Plan) error
}

type PlanValidator struct{}

func NewPlanValidator() IPlanValidator {
	return &PlanValidator{}
}

func (plv *PlanValidator) PlanValidate(plan model.Plan) error {
	return validation.ValidateStruct(&plan,
		validation.Field(
			&plan.Title,
			validation.Required.Error("Title is required"),
			validation.Length(1, 20).Error("limited max 20 characters"),
		),
		validation.Field(
			&plan.Content,
			validation.Required.Error("Content is required"),
			validation.Length(1, 100).Error("limited max 100 characters"),
		),
	)
}
