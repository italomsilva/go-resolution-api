package userMain

import (
	"database/sql"
	"go-resolution-api/application/user/controller"
	"go-resolution-api/application/user/repository"
	"go-resolution-api/application/user/router"
	"go-resolution-api/application/user/usecase"

	"github.com/gin-gonic/gin"
)

func InitializeModule(databaseConnection *sql.DB, routerGin *gin.Engine){
	UserRepository := repository.NewUserRepository(databaseConnection)
	UserUseCase := usecase.NewUserUseCase(UserRepository)
	UserController := controller.NewUserController(UserUseCase)
	router.InitializeRoutes(&UserController, routerGin)
}