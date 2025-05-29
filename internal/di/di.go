package di

import (
	"database/sql"
	"go-resolution-api/internal/controller"
	"go-resolution-api/internal/infra"
	"go-resolution-api/internal/middleware"
	"go-resolution-api/internal/routes"
	ProblemUC "go-resolution-api/internal/usecase/problem"
	SolutionUC "go-resolution-api/internal/usecase/solution"
	UserUC "go-resolution-api/internal/usecase/user"
	"go-resolution-api/pkg"

	"github.com/gin-gonic/gin"
)

func InjectDependencies(databaseConnection *sql.DB, routerGin *gin.Engine) {
	// gateways
	tokenGateway := pkg.NewAuthJWTGateway()
	idGeneratorGateway := pkg.NewUUIDGateway()
	cryptorGateway := pkg.NewBcryptGateway()

	// repositories
	problemRepository := infra.NewProblemRepository(databaseConnection)
	userRepository := infra.NewUserRepository(databaseConnection)
	solutionRepository := infra.NewSolutionRepository(databaseConnection)

	//problems usecases
	createProblemUsecase := ProblemUC.NewCreateProblemUsecase(problemRepository, tokenGateway, idGeneratorGateway)
	deleteProblemUsecase := ProblemUC.NewDeleteProblemUsecase(problemRepository, tokenGateway)
	deleteAllProblemsByUserIdUsecase := ProblemUC.NewDeleteAllProblemsByUserIdUsecase(problemRepository, tokenGateway)
	getAllProblemsByUserIdUsecase := ProblemUC.NewGetAllProblemsByUserIdUsecase(problemRepository, tokenGateway)
	getAllProblemsUsecase := ProblemUC.NewGetAllProblemsUsecase(problemRepository)
	getProblemByIdUsecase := ProblemUC.NewGetProblemByIDUsecase(problemRepository)
	updateProblemUsecase := ProblemUC.NewUpdateProblemUsecase(problemRepository, tokenGateway)

	// user usecases
	createUserUsecase := UserUC.NewCreateUserUsecase(userRepository, tokenGateway, idGeneratorGateway, cryptorGateway)
	loginUsecase := UserUC.NewLoginUsecase(userRepository, tokenGateway, cryptorGateway)
	deleteUserUsecase := UserUC.NewDeleteUserUsecase(userRepository, tokenGateway, loginUsecase)
	getUserByIdUsecase := UserUC.NewGetUserByIdUsecase(userRepository)
	getAllUsersUsecase := UserUC.NewGetUsersUsecase(userRepository)
	updateUserUsecase := UserUC.NewUpdateUserUsecase(userRepository, tokenGateway)

	// solutions usecase
	createSolutionUsecase := SolutionUC.NewCreateSolutionUsecase(solutionRepository, problemRepository, tokenGateway, idGeneratorGateway)
	getAllSolutionByProblemId := SolutionUC.NewGetAllSolutionsByProblemIdUsecase(solutionRepository, problemRepository)

	problemController := controller.NewProblemController(
		tokenGateway,
		createProblemUsecase,
		deleteAllProblemsByUserIdUsecase,
		deleteProblemUsecase,
		getAllProblemsByUserIdUsecase,
		getAllProblemsUsecase,
		getProblemByIdUsecase,
		updateProblemUsecase,
	)
	userController := controller.NewUserController(
		tokenGateway,
		createUserUsecase,
		deleteUserUsecase,
		getUserByIdUsecase,
		getAllUsersUsecase,
		loginUsecase,
		updateUserUsecase,
	)
	solutionController := controller.NewSolutionController(
		tokenGateway,
		createSolutionUsecase,
		getAllSolutionByProblemId,
	)

	authMiddleware := middleware.NewAuthMiddleware(tokenGateway)
	apiKeyMiddleware := middleware.NewApiKeyMiddleware()

	routes.InitializeUserRoutes(&userController, routerGin, authMiddleware, apiKeyMiddleware)
	routes.InitializeProblemsRoutes(&problemController, routerGin, authMiddleware, apiKeyMiddleware)
	routes.InitializeSolutionRoutes(&solutionController, routerGin, authMiddleware, apiKeyMiddleware)

}
