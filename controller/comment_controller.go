package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ICommentController interface {
	CreateComment(c echo.Context) error
	GetCommentsByPlanID(c echo.Context) error
	GetMyComments(c echo.Context) error
	DeleteComment(c echo.Context) error
}

type commentController struct {
	cu usecase.ICommentUsecase
}

func NewCommentController(cu usecase.ICommentUsecase) ICommentController {
	return &commentController{cu: cu}
}

func (cc *commentController) CreateComment(c echo.Context) error {
	comment := &model.Comment{}
	if err := c.Bind(comment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// JWTトークンがある場合はユーザーIDを設定
	if user, ok := c.Get("user").(*jwt.Token); ok {
		claims := user.Claims.(jwt.MapClaims)
		userId := uint(claims["user_id"].(float64))
		comment.UserID = &userId
	}

	res, err := cc.cu.CreateComment(comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (cc *commentController) GetCommentsByPlanID(c echo.Context) error {
	planID, err := strconv.ParseUint(c.Param("planId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid plan ID"})
	}

	comments, err := cc.cu.GetCommentsByPlanID(uint(planID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, comments)
}

func (cc *commentController) GetMyComments(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	comments, err := cc.cu.GetCommentsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, comments)
}

func (cc *commentController) DeleteComment(c echo.Context) error {
	commentID, err := strconv.ParseUint(c.Param("commentId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid comment ID"})
	}

	var userID *uint
	if user, ok := c.Get("user").(*jwt.Token); ok {
		claims := user.Claims.(jwt.MapClaims)
		id := uint(claims["user_id"].(float64))
		userID = &id
	}

	if err := cc.cu.DeleteComment(uint(commentID), userID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Comment deleted successfully"})
}
