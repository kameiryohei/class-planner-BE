package main

import (
	"backend/auth"
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/usecase"
	"backend/validator"
)

func main() {
	db := db.NewDB()
	// auth
	googleAuthConfig := auth.NewGoogleAuthConfig()

	// validation
	userValidator := validator.NewUserValidator()
	postValidator := validator.NewPostValidator()
	planValidator := validator.NewPlanValidator()

	// repository
	userRepository := repository.NewUserRepository(db)
	postRepository := repository.NewPostRepository(db)
	planRepository := repository.NewPlanRepository(db)
	courseRepository := repository.NewCourseRepository(db)
	commentRepository := repository.NewCommentRepository(db)

	// usecase
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator, googleAuthConfig)
	postUsecase := usecase.NewPostUsecase(postRepository, postValidator)
	planUsecase := usecase.NewPlanUsecase(planRepository, planValidator)
	courseUsecase := usecase.NewCourseUsecase(courseRepository)
	commentUsecase := usecase.NewCommentUsecase(commentRepository)

	// controller
	userController := controller.NewUserController(userUsecase)
	postController := controller.NewTaskController(postUsecase)
	planController := controller.NewPlanController(planUsecase)
	courseController := controller.NewCourseController(courseUsecase)
	commentController := controller.NewCommentController(commentUsecase)

	// router
	e := router.NewRouter(userController, postController, planController, courseController, commentController)
	e.Logger.Fatal(e.Start(":8080"))
}
