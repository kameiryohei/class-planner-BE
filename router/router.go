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
	cc controller.ICourseController,
	ccu controller.ICommentController,
	tc controller.ITimetableController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CorsMiddleware())
	e.Use(middleware.CsrfMiddleware())

	// グループ化
	p := e.Group("/posts")
	pl := e.Group("/plans")
	c := e.Group("/courses")
	t := e.Group("/timetable")
	comments := e.Group("/comments")
	authComments := e.Group("/comments")

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
	pl.POST("/:planId/favorite", plc.ToggleFavoritePlan)
	pl.GET("/:planId/favorite/count", plc.GetFavoriteCount)

	// courseに関するエンドポイント
	c.GET("/:courseId", cc.GetAllCourses)
	c.POST("", cc.CreateCourses)
	c.PUT("/:courseId", cc.UpdateCourse)
	c.DELETE("/:courseId", cc.DeleteCourseByID)

	// コメント関連のルート（認証不要）
	comments.Use(middleware.OptionalJwtMiddleware())
	comments.POST("", ccu.CreateComment)
	comments.GET("/plan/:planId", ccu.GetCommentsByPlanID)

	// 認証が必要なコメント関連のルート
	authComments.Use(middleware.JwtMiddleware())
	authComments.GET("/me", ccu.GetMyComments)
	authComments.DELETE("/:commentId", ccu.DeleteComment)

	// timetableに関するエンドポイント
	t.Use(middleware.JwtMiddleware())
	t.POST("/upload", tc.UploadTimetable)
	t.GET("/analyses", tc.GetUserAnalyses)
	t.GET("/analysis/:analysisId", tc.GetAnalysis)
	t.POST("/analysis/:analysisId/confirm", tc.ConfirmClasses)
	t.DELETE("/analysis/:analysisId", tc.DeleteAnalysis)

	return e
}
