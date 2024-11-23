package router

import (
	"backend/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func jwtMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	})
}

func NewRouter(uc controller.IUserController, pc controller.IPostController, plc controller.IPlanController) *echo.Echo {
	e := echo.New()

	// 環境に応じて許可するオリジンを切り替える
	var allowedOrigins []string
	if os.Getenv("GO_ENV") == "dev" {
		// 開発環境
		allowedOrigins = []string{"http://localhost:3000", os.Getenv("FE_URL")}
	} else {
		// 本番環境
		allowedOrigins = []string{os.Getenv("FE_URL")}
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode, //postmanでのテストのため
		// CookieMaxAge:   60,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CsrfToken)
	p := e.Group("/posts")
	pl := e.Group("/plans")
	// middlewareを追加
	p.Use(jwtMiddleware())

	p.GET("", pc.GetAllPosts)
	p.GET("/:planId", pc.GetPostByID)
	p.POST("", pc.CreatePost)
	p.DELETE("/:postId", pc.DeletePostByID)

	pl.Use(jwtMiddleware())
	pl.GET("", plc.GetAllPlans)
	pl.GET("/:planId", plc.GetPlansByID)
	pl.POST("", plc.CreatePlan)
	pl.DELETE("/:planId", plc.DeletePlanByID)

	return e
}
