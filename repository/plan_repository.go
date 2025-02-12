package repository

import (
	"backend/model"
	"errors"

	"gorm.io/gorm"
)

type IPlanRepository interface {
	GetAllPlans(plans *[]model.Plan, offset int, limit int) error
	GetPlanByID(plan *model.Plan, planId uint) error
	CreatePlan(plan *model.Plan) error
	UpdatePlan(plan *model.Plan, planId uint) error
	DeletePlanByID(planId uint) error
	ToggleFavoritePlan(userId uint, planId uint) error
	GetFavoriteCount(planId uint) (int64, error)
}

type planRepository struct {
	db *gorm.DB
}

func NewPlanRepository(db *gorm.DB) IPlanRepository {
	return &planRepository{db: db}
}

func (pr *planRepository) GetAllPlans(plans *[]model.Plan, offset int, limit int) error {
	return pr.db.Preload("User").
		Preload("User.University").
		Preload("User.Faculty").
		Preload("User.Department").
		Offset(offset).
		Limit(limit).
		Find(plans).Error
}

func (pr *planRepository) GetPlanByID(plan *model.Plan, planId uint) error {
	return pr.db.Preload("User").
		Preload("User.University").
		Preload("User.Faculty").
		Preload("User.Department").
		Preload("Courses").
		Preload("Posts").
		Preload("Favorites").
		Where("id = ?", planId).
		First(plan).Error
}

func (pr *planRepository) CreatePlan(plan *model.Plan) error {
	return pr.db.Create(plan).Error
}

func (pr *planRepository) UpdatePlan(plan *model.Plan, planId uint) error {
	if err := pr.db.Model(&model.Plan{}).
		Where("id = ?", planId).
		Updates(plan).Error; err != nil {
		return err
	}
	return pr.db.Where("id = ?", planId).First(plan).Error
}

func (pr *planRepository) DeletePlanByID(planId uint) error {
	result := pr.db.Where("id = ?", planId).Delete(&model.Plan{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (pr *planRepository) ToggleFavoritePlan(userId uint, planId uint) error {
	var favoritePlan model.FavoritePlan
	result := pr.db.Where("user_id = ? AND plan_id = ?", userId, planId).First(&favoritePlan)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// お気に入りが存在しない場合は新規作成
		newFavorite := model.FavoritePlan{
			UserID: userId,
			PlanID: planId,
		}
		return pr.db.Create(&newFavorite).Error
	} else if result.Error != nil {
		return result.Error
	}

	// お気に入りが既に存在する場合は削除
	return pr.db.Delete(&favoritePlan).Error
}

func (pr *planRepository) GetFavoriteCount(planId uint) (int64, error) {
	var count int64
	err := pr.db.Model(&model.FavoritePlan{}).
		Where("plan_id = ?", planId).
		Count(&count).Error
	return count, err
}
