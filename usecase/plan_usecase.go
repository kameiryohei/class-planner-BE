package usecase

import (
	"backend/model"
	"backend/repository"
	"backend/validator"
	"errors"
)

type IPlanUsecase interface {
	GetAllPlans(offset int, limit int) ([]model.PlanResponse, error)
	GetPlanByID(planId uint) (model.PlanDetailResponse, error)
	CreatePlan(plan *model.Plan) (model.PlanBaseResponse, error)
	UpdatePlan(plan *model.Plan, planId uint) (model.PlanUpdateResponse, error) // planId の型を uint に変更
	DeletePlanByID(planId uint) error
	ToggleFavoritePlan(userId, planId uint) error
	GetFavoriteCount(planId uint) (int64, error)
}

type planUsecase struct {
	pr  repository.IPlanRepository
	plv validator.IPlanValidator
}

func NewPlanUsecase(pr repository.IPlanRepository, plv validator.IPlanValidator) IPlanUsecase {
	return &planUsecase{pr: pr, plv: plv}
}

func (pu *planUsecase) GetAllPlans(offset int, limit int) ([]model.PlanResponse, error) {
	var plans []model.Plan
	if err := pu.pr.GetAllPlans(&plans, offset, limit); err != nil {
		return nil, err
	}

	// スライスの容量を事前に確保
	resPlans := make([]model.PlanResponse, 0, len(plans))
	for _, plan := range plans {
		resPlans = append(resPlans, model.PlanResponse{
			ID:        plan.ID,
			Title:     plan.Title,
			Content:   plan.Content,
			UserID:    plan.UserID,
			CreatedAt: plan.CreatedAt,
			UpdatedAt: plan.UpdatedAt,
			UserResponse: model.UserResponse{
				ID:         plan.User.ID,
				Email:      plan.User.Email,
				University: plan.User.University,
				Faculty:    plan.User.Faculty,
				Department: plan.User.Department,
			},
		})
	}
	return resPlans, nil
}

func (pu *planUsecase) GetPlanByID(planId uint) (model.PlanDetailResponse, error) {
	var plan model.Plan
	if err := pu.pr.GetPlanByID(&plan, planId); err != nil {
		return model.PlanDetailResponse{}, err
	}

	courses := make([]model.CourseResponse, 0, len(plan.Courses))
	for _, course := range plan.Courses {
		courses = append(courses, model.CourseResponse{
			ID:      course.ID,
			Name:    course.Name,
			Content: course.Content,
		})
	}

	posts := make([]model.PostResponse, 0, len(plan.Posts))
	for _, post := range plan.Posts {
		posts = append(posts, model.PostResponse{
			ID:        post.ID,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
		})
	}

	favorites := make([]model.FavoritePlanResponse, 0, len(plan.Favorites))
	for _, favorite := range plan.Favorites {
		favorites = append(favorites, model.FavoritePlanResponse{
			ID:     favorite.ID,
			UserID: favorite.UserID,
			PlanID: favorite.PlanID,
		})
	}

	resPlan := model.PlanDetailResponse{
		ID:        plan.ID,
		Title:     plan.Title,
		Content:   plan.Content,
		UserID:    plan.UserID,
		CreatedAt: plan.CreatedAt,
		UpdatedAt: plan.UpdatedAt,
		User: model.UserResponse{
			ID:         plan.User.ID,
			Email:      plan.User.Email,
			University: plan.User.University,
			Faculty:    plan.User.Faculty,
			Department: plan.User.Department,
		},
		Courses:   courses,
		Posts:     posts,
		Favorites: favorites,
	}
	return resPlan, nil
}

func (pu *planUsecase) CreatePlan(plan *model.Plan) (model.PlanBaseResponse, error) {
	// nilチェック（必要に応じて）
	if plan == nil {
		return model.PlanBaseResponse{}, errors.New("plan is nil")
	}
	if err := pu.plv.PlanValidate(*plan); err != nil {
		return model.PlanBaseResponse{}, err
	}
	if err := pu.pr.CreatePlan(plan); err != nil {
		return model.PlanBaseResponse{}, err
	}
	resPlan := model.PlanBaseResponse{
		ID:        plan.ID,
		Title:     plan.Title,
		Content:   plan.Content,
		UserID:    plan.UserID,
		CreatedAt: plan.CreatedAt,
		UpdatedAt: plan.UpdatedAt,
	}
	return resPlan, nil
}

func (pu *planUsecase) UpdatePlan(plan *model.Plan, planId uint) (model.PlanUpdateResponse, error) {
	// nilチェック
	if plan == nil {
		return model.PlanUpdateResponse{}, errors.New("plan is nil")
	}
	if err := pu.plv.PlanValidate(*plan); err != nil {
		return model.PlanUpdateResponse{}, err
	}
	if err := pu.pr.UpdatePlan(plan, planId); err != nil {
		return model.PlanUpdateResponse{}, err
	}
	resPlan := model.PlanUpdateResponse{
		ID:        plan.ID,
		Title:     plan.Title,
		Content:   plan.Content,
		CreatedAt: plan.CreatedAt,
		UpdatedAt: plan.UpdatedAt,
	}
	return resPlan, nil
}

func (pu *planUsecase) DeletePlanByID(planId uint) error {
	return pu.pr.DeletePlanByID(planId)
}

func (pu *planUsecase) ToggleFavoritePlan(userId, planId uint) error {
	return pu.pr.ToggleFavoritePlan(userId, planId)
}

func (pu *planUsecase) GetFavoriteCount(planId uint) (int64, error) {
	return pu.pr.GetFavoriteCount(planId)
}
