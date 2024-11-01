package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IPostController interface {
	GetAllPosts(c echo.Context) error
	GetPostByID(c echo.Context) error
	CreatePost(c echo.Context) error
	DeletePostByID(c echo.Context) error
}

type postController struct {
	pu usecase.IPostUsecase
}

func NewTaskController(pu usecase.IPostUsecase) IPostController {
	return &postController{pu}
}

func (pc *postController) GetAllPosts(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	postRes, err := pc.pu.GetAllPosts(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postRes)
}

func (pc *postController) GetPostByID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("postId")
	postId, _ := strconv.Atoi(id)
	postRes, err := pc.pu.GetPostByID(uint(userId.(float64)), uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postRes)
}

func (pc *postController) CreatePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	post := model.Post{}
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	post.AuthorID = uint(userId.(float64))
	postRes, err := pc.pu.CreatePost(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, postRes)
}

func (pc *postController) DeletePostByID(c echo.Context) error {
	id := c.Param("postId")
	postId, _ := strconv.Atoi(id)

	err := pc.pu.DeletePostByID(uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
