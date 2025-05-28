package di

import (
	"database/sql"
	"go-resolution-api/internal/controller"
	"go-resolution-api/internal/infra"
	"go-resolution-api/internal/middleware"
	"go-resolution-api/internal/routes"
	ProblemUC "go-resolution-api/internal/usecase/problem"
	UserUC "go-resolution-api/internal/usecase/user"
	"go-resolution-api/pkg"

	"github.com/gin-gonic/gin"
)

func InjectDependencies(databaseConnection *sql.DB, routerGin *gin.Engine) {
	tokenGateway := pkg.NewAuthJWTGateway()
	idGeneratorGateway := pkg.NewUUIDGateway()
	cryptorGateway := pkg.NewBcryptGateway()

	problemRepository := infra.NewProblemRepository(databaseConnection)
	UserRepository := infra.NewUserRepository(databaseConnection)

	problemUseCase := ProblemUC.NewProblemUseCase(problemRepository, tokenGateway, idGeneratorGateway)
	UserUseCase := UserUC.NewUserUseCase(UserRepository, tokenGateway, cryptorGateway, idGeneratorGateway)

	problemController := controller.NewProblemController(problemUseCase, tokenGateway)
	UserController := controller.NewUserController(UserUseCase, tokenGateway)

	authMiddleware := middleware.NewAuthMiddleware(tokenGateway)
	apiKeyMiddleware := middleware.NewApiKeyMiddleware()

	routes.InitializeUserRoutes(&UserController, routerGin, authMiddleware, apiKeyMiddleware)
	routes.InitializeProblemsRoutes(&problemController, routerGin, authMiddleware, apiKeyMiddleware)

}
