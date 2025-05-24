package userMain

import (
	"database/sql"
	"go-resolution-api/application/user/controller"
	"go-resolution-api/application/user/repository"
	"go-resolution-api/application/user/router"
	"go-resolution-api/application/user/usecase"
)

func InitializeModule(databaseConnection *sql.DB){
	UserRepository := repository.NewUserRepository(databaseConnection)
	UserUseCase := usecase.NewUserUseCase(UserRepository)
	UserController := controller.NewUserController(UserUseCase)
	router.InitializeRoutes(&UserController)
}