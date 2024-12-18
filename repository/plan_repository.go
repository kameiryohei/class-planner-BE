package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type IPlanRepository interface {
	GetAllPlans(plans *[]model.Plan, offset int, limit int) error
	GetPlanByID(plan *model.Plan, planId uint) error
	CreatePlan(plan *model.Plan) error
	UpdatePlan(plan *model.Plan, planId int) error
	DeletePlanByID(planId uint) error
}

type planRepository struct {
	db *gorm.DB
}

func NewPlanRepository(db *gorm.DB) IPlanRepository {
	return &planRepository{db}
}

func (pr *planRepository) GetAllPlans(plans *[]model.Plan, offset int, limit int) error {
	if err := pr.db.Preload("User").
		Preload("User.University").
		Preload("User.Faculty").
		Preload("User.Department").
		Offset(offset).
		Limit(limit).
		Find(plans).Error; err != nil {
		return err
	}
	return nil
}

func (pr *planRepository) GetPlanByID(plan *model.Plan, planId uint) error {
	if err := pr.db.Preload("User").
		Preload("User.University").
		Preload("User.Faculty").
		Preload("User.Department").
		Preload("Courses").
		Preload("Posts").
		Preload("Favorites").
		Where("id = ?", planId).First(plan).Error; err != nil {
		return err
	}
	return nil
}

func (pr *planRepository) CreatePlan(plan *model.Plan) error {
	if err := pr.db.Create(plan).Error; err != nil {
		return err
	}
	return nil
}

func (pr *planRepository) UpdatePlan(plan *model.Plan, planId int) error {
	if err := pr.db.Model(&model.Plan{}).Where("id = ?", planId).Updates(plan).Error; err != nil {
		return err
	}
	if err := pr.db.Where("id = ?", planId).First(plan).Error; err != nil {
		return err
	}
	return nil
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
