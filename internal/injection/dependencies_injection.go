package injection

import (
	"database/sql"
	"go-resolution-api/internal/controller"
	"go-resolution-api/internal/infra"
	"go-resolution-api/internal/middleware"
	"go-resolution-api/internal/routes"
	ProblemUsecase "go-resolution-api/internal/usecase/problem"
	SolutionUsecase "go-resolution-api/internal/usecase/solution"
	UserUsecase "go-resolution-api/internal/usecase/user"
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
	createProblemUsecase := ProblemUsecase.NewCreateProblemUsecase(problemRepository, tokenGateway, idGeneratorGateway)
	deleteProblemUsecase := ProblemUsecase.NewDeleteProblemUsecase(problemRepository, tokenGateway)
	deleteAllProblemsByUserIdUsecase := ProblemUsecase.NewDeleteAllProblemsByUserIdUsecase(problemRepository, tokenGateway)
	getAllProblemsByUserIdUsecase := ProblemUsecase.NewGetAllProblemsByUserIdUsecase(problemRepository, tokenGateway)
	getAllProblemsUsecase := ProblemUsecase.NewGetAllProblemsUsecase(problemRepository)
	getProblemByIdUsecase := ProblemUsecase.NewGetProblemByIDUsecase(problemRepository)
	updateProblemUsecase := ProblemUsecase.NewUpdateProblemUsecase(problemRepository, tokenGateway)

	// user usecases
	createUserUsecase := UserUsecase.NewCreateUserUsecase(userRepository, tokenGateway, idGeneratorGateway, cryptorGateway)
	loginUsecase := UserUsecase.NewLoginUsecase(userRepository, tokenGateway, cryptorGateway)
	deleteUserUsecase := UserUsecase.NewDeleteUserUsecase(userRepository, tokenGateway, loginUsecase)
	getUserByIdUsecase := UserUsecase.NewGetUserByIdUsecase(userRepository)
	getAllUsersUsecase := UserUsecase.NewGetUsersUsecase(userRepository)
	updateUserUsecase := UserUsecase.NewUpdateUserUsecase(userRepository, tokenGateway)

	// solutions usecase
	createSolutionUsecase := SolutionUsecase.NewCreateSolutionUsecase(solutionRepository, problemRepository, tokenGateway, idGeneratorGateway)
	getAllSolutionsByProblemIdUsecase := SolutionUsecase.NewGetAllSolutionsByProblemIdUsecase(solutionRepository, problemRepository)
	getSolutionByIdUsecase := SolutionUsecase.NewGetSolutionByIdUsecase(solutionRepository)
	deleteSolutionUsecase := SolutionUsecase.NewDeleteSolutionUsecase(solutionRepository, tokenGateway)
	deleteAllSolutionsByProblemIdUsecase := SolutionUsecase.NewDeleteAllSolutionsByProblemIdUsecase(solutionRepository, problemRepository, tokenGateway)
	deleteAllSolutionsByUserIdUsecase := SolutionUsecase.NewDeleteAllSolutionsByUserIdUsecase(solutionRepository, userRepository)
	updateSolutionUsecase := SolutionUsecase.NewUpdateSolutionUsecase(solutionRepository, tokenGateway)

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
		getAllSolutionsByProblemIdUsecase,
		getSolutionByIdUsecase,
		deleteSolutionUsecase,
		deleteAllSolutionsByProblemIdUsecase,
		deleteAllSolutionsByUserIdUsecase,
		updateSolutionUsecase,
	)

	authMiddleware := middleware.NewAuthMiddleware(tokenGateway)
	apiKeyMiddleware := middleware.NewApiKeyMiddleware()

	routes.InitializeUserRoutes(&userController, routerGin, authMiddleware, apiKeyMiddleware)
	routes.InitializeProblemsRoutes(&problemController, routerGin, authMiddleware, apiKeyMiddleware)
	routes.InitializeSolutionRoutes(&solutionController, routerGin, authMiddleware, apiKeyMiddleware)

}
