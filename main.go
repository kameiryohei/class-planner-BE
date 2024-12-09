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
	googleAuthConfig := auth.NewGoogleAuthConfig()
	userValidator := validator.NewUserValidator()
	postValidator := validator.NewPostValidator()
	planValidator := validator.NewPlanValidator()
	userRepository := repository.NewUserRepository(db)
	postRepository := repository.NewPostRepository(db)
	planRepository := repository.NewPlanRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator, googleAuthConfig)
	postUsecase := usecase.NewPostUsecase(postRepository, postValidator)
	planUsecase := usecase.NewPlanUsecase(planRepository, planValidator)
	userController := controller.NewUserController(userUsecase)
	postController := controller.NewTaskController(postUsecase)
	planController := controller.NewPlanController(planUsecase)
	e := router.NewRouter(userController, postController, planController)
	e.Logger.Fatal(e.Start(":8080"))
}
