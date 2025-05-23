package main

import (
	"go-resolution-api/controller"
	"go-resolution-api/database"
	"go-resolution-api/repository"
	"go-resolution-api/router"
	"go-resolution-api/usecase"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	databaseConnection := database.ConnectDB()
	if databaseConnection == nil {
		panic("Error opening connection to the database")
	}
	
	UserRepository := repository.NewUserRepository(databaseConnection)
	UserUseCase := usecase.NewUserUseCase(UserRepository)
	UserController := controller.NewUserController(UserUseCase)

	router.InitializeRoutes(&UserController)
	
}
