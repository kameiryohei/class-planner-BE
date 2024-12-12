package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type ICourseRepository interface {
	GetAllCourses(courses *[]model.Course, planId uint) error
	CreateCourse(course *[]model.Course) error
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

func (cr *courseRepository) CreateCourse(course *[]model.Course) error {
	for _, c := range *course {
		if err := cr.db.Create(&c).Error; err != nil {
			return err
		}
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
	if err := cr.db.Where("id = ?", courseId).Delete(&model.Course{}).Error; err != nil {
		return err
	}
	return nil
}
