package main

import (
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/usecase"
	"backend/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	postValidator := validator.NewPostValidator()
	userRepository := repository.NewUserRepository(db)
	postRepository := repository.NewPostRepository(db)
	planRepository := repository.NewPlanRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	postUsecase := usecase.NewPostUsecase(postRepository, postValidator)
	planUsecase := usecase.NewPlanUsecase(planRepository)
	userController := controller.NewUserController(userUsecase)
	postController := controller.NewTaskController(postUsecase)
	planController := controller.NewPlanController(planUsecase)
	e := router.NewRouter(userController, postController, planController)
	e.Logger.Fatal(e.Start(":8080"))
}
