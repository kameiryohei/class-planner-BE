package router

import (
	"backend/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, pc controller.IPostController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.Logout)
	t := e.Group("/posts")
	// middlewareを追加
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", pc.GetAllPosts)
	t.GET("/:postId", pc.GetPostByID)
	t.POST("", pc.CreatePost)
	t.DELETE("/:postId", pc.DeletePostByID)
	return e
}
