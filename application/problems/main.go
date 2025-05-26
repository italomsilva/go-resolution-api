package problemMain

import (
	"database/sql"
	"go-resolution-api/application/problems/controller"
	"go-resolution-api/application/problems/repository"
	"go-resolution-api/application/problems/router"
	"go-resolution-api/application/problems/usecase"

	"github.com/gin-gonic/gin"
)

func InitializeModule(databaseConnection *sql.DB, routerGin *gin.Engine) {
	problemRepository := repository.NewProblemRepository(databaseConnection)
	problemUseCase := usecase.NewProblemUseCase(&problemRepository)
	problemController := controller.NewProblemController(&problemUseCase)
	router.InitializeRoutes(&problemController, routerGin)
}