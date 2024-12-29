package router

import (
	"backend/controller"
	"backend/middleware"

	"github.com/labstack/echo/v4"
)

func NewRouter(
	uc controller.IUserController,
	pc controller.IPostController,
	plc controller.IPlanController,
	cc controller.ICourseController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CorsMiddleware())
	e.Use(middleware.CsrfMiddleware())

	// グループ化
	p := e.Group("/posts")
	pl := e.Group("/plans")
	c := e.Group("/courses")

	// 認証に関するエンドポイント
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CsrfToken)
	e.GET("/auth/google/login", uc.GoogleLogin)
	e.GET("/auth/google/callback", uc.GoogleCallback)

	// postに関するエンドポイント
	p.Use(middleware.JwtMiddleware())
	p.GET("", pc.GetAllPosts)
	p.GET("/:planId", pc.GetPostByID)
	p.POST("", pc.CreatePost)
	p.DELETE("/:postId", pc.DeletePostByID)

	// planに関するエンドポイント
	pl.Use(middleware.JwtMiddleware())
	pl.GET("", plc.GetAllPlans)
	pl.GET("/:planId", plc.GetPlansByID)
	pl.POST("", plc.CreatePlan)
	pl.PUT("/:planId", plc.UpdatePlan)
	pl.DELETE("/:planId", plc.DeletePlanByID)

	// courseに関するエンドポイント
	c.GET("/:courseId", cc.GetAllCourses)
	c.POST("", cc.CreateCourses)
	c.PUT("/:courseId", cc.UpdateCourse)
	c.DELETE("/:courseId", cc.DeleteCourseByID)

	return e
}
