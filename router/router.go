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
	p := e.Group("/posts")
	// middlewareを追加
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	p.GET("", pc.GetAllPosts)
	p.GET("/:postId", pc.GetPostByID)
	p.POST("", pc.CreatePost)
	p.DELETE("/:postId", pc.DeletePostByID)
	return e
}
