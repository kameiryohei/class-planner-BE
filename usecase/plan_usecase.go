package usecase

import (
	"backend/model"
	"backend/repository"
)

type IPlanUsecase interface {
	GetAllPlans() ([]model.PlanResponse, error)
	GetPlanByID(planId uint) (model.Plan, error)
	CreatePlan(plan *model.Plan) (model.PlanResponse, error)
	UpdatePlan(plan *model.Plan) (model.PlanResponse, error)
	DeletePlanByID(planId uint) error
}

type planUsecase struct {
	pr repository.IPlanRepository
}

func NewPlanUsecase(pr repository.IPlanRepository) IPlanUsecase {
	return &planUsecase{pr}
}

func (pu *planUsecase) GetAllPlans() ([]model.PlanResponse, error) {
	plans := []model.Plan{}
	if err := pu.pr.GetAllPlans(&plans); err != nil {
		return nil, err
	}
	resPlans := []model.PlanResponse{}
	for _, v := range plans {
		p := model.PlanResponse{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			UserID:    v.UserID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User:      v.User,
		}
		resPlans = append(resPlans, p)
	}
	return resPlans, nil
}

func (pu *planUsecase) GetPlanByID(planId uint) (model.Plan, error) {
	plan := model.Plan{}
	if err := pu.pr.GetPlanByID(&plan, planId); err != nil {
		return model.Plan{}, err
	}
	resPlan := model.Plan{
		ID:        plan.ID,
		Title:     plan.Title,
		Content:   plan.Content,
		UserID:    plan.UserID,
		CreatedAt: plan.CreatedAt,
		UpdatedAt: plan.UpdatedAt,
		User:      plan.User,
		Courses:   plan.Courses,
		Posts:     plan.Posts,
		Favorites: plan.Favorites,
	}
	return resPlan, nil
}

func (pu *planUsecase) CreatePlan(plan *model.Plan) (model.PlanResponse, error) {
	if err := pu.pr.CreatePlan(plan); err != nil {
		return model.PlanResponse{}, err
	}
	resPlan := model.PlanResponse{
		ID:        plan.ID,
		Title:     plan.Title,
		Content:   plan.Content,
		UserID:    plan.UserID,
		CreatedAt: plan.CreatedAt,
		UpdatedAt: plan.UpdatedAt,
		User:      plan.User,
	}
	return resPlan, nil
}

func (pu *planUsecase) UpdatePlan(plan *model.Plan) (model.PlanResponse, error) {
	if err := pu.pr.UpdatePlan(plan); err != nil {
		return model.PlanResponse{}, err
	}
	resPlan := model.PlanResponse{
		ID:        plan.ID,
		Title:     plan.Title,
		Content:   plan.Content,
		UserID:    plan.UserID,
		CreatedAt: plan.CreatedAt,
		UpdatedAt: plan.UpdatedAt,
		User:      plan.User,
	}
	return resPlan, nil
}

func (pu *planUsecase) DeletePlanByID(planId uint) error {
	if err := pu.pr.DeletePlanByID(planId); err != nil {
		return err
	}
	return nil
}
