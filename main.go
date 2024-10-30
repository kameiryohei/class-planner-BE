package main

import (
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/usecase"
)

func main (){
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db) 
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e:= router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}