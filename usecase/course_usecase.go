package usecase

import (
	"backend/model"
	"backend/repository"
)

type ICourseUsecase interface {
	GetAllCourses(planId uint) ([]model.CourseResponse, error)
	CreateCourses(courses []model.Course) ([]model.CourseResponse, error)
	UpdateCourse(course *model.Course, courseId int) (model.CourseResponse, error)
	DeleteCourseByID(courseId uint) error
}

type courseUsecase struct {
	cr repository.ICourseRepository
}

func NewCourseUsecase(cr repository.ICourseRepository) ICourseUsecase {
	return &courseUsecase{cr}
}

func (cu *courseUsecase) GetAllCourses(planId uint) ([]model.CourseResponse, error) {
	courses := []model.Course{}
	if err := cu.cr.GetAllCourses(&courses, planId); err != nil {
		return nil, err
	}
	resCourses := []model.CourseResponse{}
	for _, v := range courses {
		c := model.CourseResponse{
			ID:      v.ID,
			Name:    v.Name,
			Content: v.Content,
		}
		resCourses = append(resCourses, c)
	}
	return resCourses, nil
}

func (cu *courseUsecase) CreateCourses(courses []model.Course) ([]model.CourseResponse, error) {
	if err := cu.cr.CreateCourses(&courses); err != nil {
		return nil, err
	}
	resCourses := []model.CourseResponse{}
	for _, v := range courses {
		c := model.CourseResponse{
			ID:      v.ID,
			Name:    v.Name,
			Content: v.Content,
		}
		resCourses = append(resCourses, c)
	}
	return resCourses, nil
}

func (cu *courseUsecase) UpdateCourse(course *model.Course, courseId int) (model.CourseResponse, error) {
	if err := cu.cr.UpdateCourse(course, courseId); err != nil {
		return model.CourseResponse{}, err
	}
	return model.CourseResponse{}, nil
}

func (cu *courseUsecase) DeleteCourseByID(courseId uint) error {
	if err := cu.cr.DeleteCourseByID(courseId); err != nil {
		return err
	}
	return nil
}
