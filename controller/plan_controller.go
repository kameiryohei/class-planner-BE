package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IPlanController interface {
	GetAllPlans(c echo.Context) error
	GetPlansByID(c echo.Context) error
	CreatePlan(c echo.Context) error
	UpdatePlan(c echo.Context) error
	DeletePlanByID(c echo.Context) error
	ToggleFavoritePlan(c echo.Context) error
	GetFavoriteCount(c echo.Context) error
}

type planController struct {
	pu usecase.IPlanUsecase
}

func NewPlanController(pu usecase.IPlanUsecase) IPlanController {
	return &planController{pu}
}

func (pc *planController) GetAllPlans(c echo.Context) error {
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10 // デフォルトのリミットをは10に設定
	}
	postRes, err := pc.pu.GetAllPlans(offset, limit)
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	plan := &model.Plan{}
	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	plan.UserID = userId
	planRes, err := pc.pu.CreatePlan(plan)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, planRes)
}

func (pc *planController) UpdatePlan(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	id := c.Param("planId")
	planId, _ := strconv.Atoi(id)

	plan := &model.Plan{}
	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	plan.UserID = userId
	planRes, err := pc.pu.UpdatePlan(plan, uint(planId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
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

func (pc *planController) ToggleFavoritePlan(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	id := c.Param("planId")
	planId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid plan ID"})
	}

	if err := pc.pu.ToggleFavoritePlan(userId, uint(planId)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Favorite toggled successfully"})
}

func (pc *planController) GetFavoriteCount(c echo.Context) error {
	id := c.Param("planId")
	planId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid plan ID"})
	}

	count, err := pc.pu.GetFavoriteCount(uint(planId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"plan_id":        planId,
		"favorite_count": count,
	})
}
