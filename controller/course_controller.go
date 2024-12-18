package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ICourseController interface {
	GetAllCourses(c echo.Context) error
	CreateCourses(c echo.Context) error
	UpdateCourse(c echo.Context) error
	DeleteCourseByID(c echo.Context) error
}

type courseController struct {
	cu usecase.ICourseUsecase
}

func NewCourseController(cu usecase.ICourseUsecase) ICourseController {
	return &courseController{cu}
}

func (cc *courseController) GetAllCourses(c echo.Context) error {
	id := c.Param("courseId")
	planId, _ := strconv.Atoi(id)
	postRes, err := cc.cu.GetAllCourses(uint(planId))
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, postRes)
}

func (cc *courseController) CreateCourses(c echo.Context) error {
	var courses []model.Course
	if err := c.Bind(&courses); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdCourses, err := cc.cu.CreateCourses(courses)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdCourses)
}

func (cc *courseController) UpdateCourse(c echo.Context) error {
	id := c.Param("courseId")
	courseId, _ := strconv.Atoi(id)

	course := &model.Course{}
	if err := c.Bind(course); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	postRes, err := cc.cu.UpdateCourse(course, courseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postRes)
}

func (cc *courseController) DeleteCourseByID(c echo.Context) error {
	id := c.Param("courseId")
	courseId, _ := strconv.Atoi(id)

	err := cc.cu.DeleteCourseByID(uint(courseId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Deleted")
}
