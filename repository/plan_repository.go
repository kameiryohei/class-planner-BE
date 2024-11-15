package repository

import (
	"backend/model"
	"fmt"

	"gorm.io/gorm"
)

type IPlanRepository interface {
	GetAllPlans(plans *[]model.Plan) error
	GetPlanByID(plan *model.PlanDetailResponse, planId uint) error
	CreatePlan(plan *model.Plan) error
	UpdatePlan(plan *model.Plan) error
	DeletePlanByID(planId uint) error
}

type planRepository struct {
	db *gorm.DB
}

func NewPlanRepository(db *gorm.DB) IPlanRepository {
	return &planRepository{db}
}

func (pr *planRepository) GetAllPlans(plans *[]model.Plan) error {
	if err := pr.db.Preload("User").Preload("Courses").Preload("Posts").Preload("Favorites").Find(plans).Error; err != nil {
		return err
	}
	fmt.Println(plans)
	return nil
}

func (pr *planRepository) GetPlanByID(plan *model.PlanDetailResponse, planId uint) error {
	if err := pr.db.Preload("User").Preload("Courses").Preload("Posts").Preload("Favorites").Where("id = ?", planId).First(plan).Error; err != nil {
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

func (pr *planRepository) UpdatePlan(plan *model.Plan) error {
	if err := pr.db.Save(plan).Error; err != nil {
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
