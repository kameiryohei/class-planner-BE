package main

import (
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	postRepository := repository.NewPostRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	postUsecase := usecase.NewPostUsecase(postRepository)
	userController := controller.NewUserController(userUsecase)
	postController := controller.NewTaskController(postUsecase)
	e := router.NewRouter(userController, postController)
	e.Logger.Fatal(e.Start(":8080"))
}
