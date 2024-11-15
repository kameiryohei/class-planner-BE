package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IPlanController interface {
	GetAllPlans(c echo.Context) error
	GetPlansByID(c echo.Context) error
	CreatePlan(c echo.Context) error
	DeletePlanByID(c echo.Context) error
}

type planController struct {
	pu usecase.IPlanUsecase
}

func NewPlanController(pu usecase.IPlanUsecase) IPlanController {
	return &planController{pu}
}

func (pc *planController) GetAllPlans(c echo.Context) error {
	postRes, err := pc.pu.GetAllPlans()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postRes)
}

func (pc *planController) GetPlansByID(c echo.Context) error {
	id := c.Param("planId")
	planId, _ := strconv.Atoi(id)
	postRes, err := pc.pu.GetPlanByID(uint(planId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postRes)
}

func (pc *planController) CreatePlan(c echo.Context) error {
	plan := &model.Plan{}
	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	planRes, err := pc.pu.CreatePlan(plan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, planRes)
}

func (pc *planController) DeletePlanByID(c echo.Context) error {
	id := c.Param("planId")
	planId, _ := strconv.Atoi(id)
	err := pc.pu.DeletePlanByID(uint(planId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Plan Deleted")
}
