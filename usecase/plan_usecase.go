package usecase

import (
	"backend/model"
	"backend/repository"
)

type IPlanUsecase interface {
	GetAllPlans(offset int, limit int) ([]model.PlanResponse, error)
	GetPlanByID(planId uint) (model.PlanDetailResponse, error)
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

func (pu *planUsecase) GetAllPlans(offset int, limit int) ([]model.PlanResponse, error) {
	plans := []model.Plan{}
	if err := pu.pr.GetAllPlans(&plans, offset, limit); err != nil {
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
			UserResponse: model.UserResponse{
				ID:         v.User.ID,
				Email:      v.User.Email,
				University: v.User.University,
				Faculty:    v.User.Faculty,
				Department: v.User.Department,
			},
		}
		resPlans = append(resPlans, p)
	}
	return resPlans, nil
}

func (pu *planUsecase) GetPlanByID(planId uint) (model.PlanDetailResponse, error) {
	plan := model.Plan{}
	if err := pu.pr.GetPlanByID(&plan, planId); err != nil {
		return model.PlanDetailResponse{}, err
	}

	courses := []model.CourseResponse{}
	for _, course := range plan.Courses {
		courses = append(courses, model.CourseResponse{
			ID:      course.ID,
			Name:    course.Name,
			Content: course.Content,
		})
	}
	posts := []model.PostResponse{}
	for _, post := range plan.Posts {
		posts = append(posts, model.PostResponse{
			ID:        post.ID,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
		})
	}
	favorites := []model.FavoritePlanResponse{}
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
	}
	return resPlan, nil
}

func (pu *planUsecase) DeletePlanByID(planId uint) error {
	if err := pu.pr.DeletePlanByID(planId); err != nil {
		return err
	}
	return nil
}
