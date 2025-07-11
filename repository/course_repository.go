package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type ICourseRepository interface {
	GetAllCourses(courses *[]model.Course, planId uint) error
	CreateCourse(course *model.Course) error
	CreateCourses(courses *[]model.Course) error
	UpdateCourse(course *model.Course, courseId int) error
	DeleteCourseByID(courseId uint) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) ICourseRepository {
	return &courseRepository{db}
}

func (cr *courseRepository) GetAllCourses(courses *[]model.Course, planId uint) error {
	if err := cr.db.Where("plan_id = ?", planId).Find(courses).Error; err != nil {
		return err
	}
	return nil
}

func (cr *courseRepository) CreateCourse(course *model.Course) error {
	if err := cr.db.Create(course).Error; err != nil {
		return err
	}
	return nil
}

func (cr *courseRepository) CreateCourses(courses *[]model.Course) error {
	if err := cr.db.Create(courses).Error; err != nil {
		return err
	}
	return nil
}

func (cr *courseRepository) UpdateCourse(course *model.Course, courseId int) error {
	if err := cr.db.Model(&model.Course{}).Where("id = ?", courseId).Updates(course).Error; err != nil {
		return err
	}
	return nil
}

func (cr *courseRepository) DeleteCourseByID(courseId uint) error {
	result := cr.db.Where("id = ?", courseId).Delete(&model.Course{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
