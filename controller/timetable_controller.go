package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITimetableController interface {
	UploadTimetable(c echo.Context) error
	GetAnalysis(c echo.Context) error
	GetUserAnalyses(c echo.Context) error
	ConfirmClasses(c echo.Context) error
	DeleteAnalysis(c echo.Context) error
}

type timetableController struct {
	tu usecase.ITimetableUsecase
}

func NewTimetableController(tu usecase.ITimetableUsecase) ITimetableController {
	return &timetableController{tu}
}

func (tc *timetableController) UploadTimetable(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	req := &model.TimetableUploadRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	analysisResponse, err := tc.tu.UploadAndAnalyze(*req, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Timetable uploaded and analyzed successfully",
		"data":    analysisResponse,
	})
}

func (tc *timetableController) GetAnalysis(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	analysisIDStr := c.Param("analysisId")
	analysisID, err := strconv.ParseUint(analysisIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid analysis ID",
		})
	}

	analysisResponse, err := tc.tu.GetAnalysis(uint(analysisID), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": analysisResponse,
	})
}

func (tc *timetableController) GetUserAnalyses(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	analysesResponse, err := tc.tu.GetUserAnalyses(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": analysesResponse,
	})
}

func (tc *timetableController) ConfirmClasses(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	analysisIDStr := c.Param("analysisId")
	analysisID, err := strconv.ParseUint(analysisIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid analysis ID",
		})
	}

	req := &model.ConfirmClassesRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	if err := tc.tu.ConfirmClasses(uint(analysisID), *req, userID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Classes confirmed and added to plan successfully",
	})
}

func (tc *timetableController) DeleteAnalysis(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	analysisIDStr := c.Param("analysisId")
	analysisID, err := strconv.ParseUint(analysisIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid analysis ID",
		})
	}

	if err := tc.tu.DeleteAnalysis(uint(analysisID), userID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Analysis deleted successfully",
	})
}