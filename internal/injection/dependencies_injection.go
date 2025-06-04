package injection

import (
	"database/sql"
	"go-resolution-api/internal/controller"
	infraGateway "go-resolution-api/internal/infra/gateway"
	infraRepository "go-resolution-api/internal/infra/repository"
	"go-resolution-api/internal/middleware"
	"go-resolution-api/internal/routes"
	ProblemUsecase "go-resolution-api/internal/usecase/problem"
	ProblemSectorUsecase "go-resolution-api/internal/usecase/problem/problem_sector"
	SectorUsecase "go-resolution-api/internal/usecase/sector"
	SolutionUsecase "go-resolution-api/internal/usecase/solution"
	SolutionReactionUsecase "go-resolution-api/internal/usecase/solution/solution_reaction"
	UserUsecase "go-resolution-api/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

func InjectDependencies(databaseConnection *sql.DB, routerGin *gin.Engine) {
	// gateways
	tokenGateway := infraGateway.NewAuthJWTGateway()
	idGeneratorGateway := infraGateway.NewUUIDGateway()
	cryptorGateway := infraGateway.NewBcryptGateway()

	// repositories
	problemRepository := infraRepository.NewProblemRepository(databaseConnection)
	userRepository := infraRepository.NewUserRepository(databaseConnection)
	solutionRepository := infraRepository.NewSolutionRepository(databaseConnection)
	sectorRepository := infraRepository.NewSectorRepository(databaseConnection)
	problemSectorRepository := infraRepository.NewProblemSectorRepository(databaseConnection)
	solutionReactionRepository := infraRepository.NewSolutionReactionRepository(databaseConnection)

	//problems usecases
	createProblemUsecase := ProblemUsecase.NewCreateProblemUsecase(problemRepository, tokenGateway, idGeneratorGateway)
	deleteProblemUsecase := ProblemUsecase.NewDeleteProblemUsecase(problemRepository, tokenGateway)
	deleteAllProblemsByUserIdUsecase := ProblemUsecase.NewDeleteAllProblemsByUserIdUsecase(problemRepository, tokenGateway)
	getAllProblemsByUserIdUsecase := ProblemUsecase.NewGetAllProblemsByUserIdUsecase(problemRepository, tokenGateway)
	getAllProblemsUsecase := ProblemUsecase.NewGetAllProblemsUsecase(problemRepository)
	getProblemByIdUsecase := ProblemUsecase.NewGetProblemByIDUsecase(problemRepository)
	updateProblemUsecase := ProblemUsecase.NewUpdateProblemUsecase(problemRepository, tokenGateway)

	//problem sectors usecases
	createProblemSectorUsecase := ProblemSectorUsecase.NewCreateProblemSectorUsecase(problemSectorRepository, problemRepository, sectorRepository, tokenGateway)
	deleteProblemSectorUsecase := ProblemSectorUsecase.NewDeleteProblemSectorUsecase(problemSectorRepository, problemRepository, tokenGateway)
	getAllSectorsByProblemIdUsecase := ProblemSectorUsecase.NewGetAllSectorsByProblemIdUsecase(problemSectorRepository, problemRepository, sectorRepository, tokenGateway)

	//sector usecase
	getAllSectorsUsecase := SectorUsecase.NewGetAllSectorsUsecase(sectorRepository)
	getSectorByIdUsecase := SectorUsecase.NewGetSectorByIDUsecase(sectorRepository)

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

	//solutions reactions usecases
	createSolutionReactionUsecase := SolutionReactionUsecase.NewCreateSolutionReactionUsecase(solutionReactionRepository, solutionRepository, tokenGateway, idGeneratorGateway)
	deleteSolutionReactionUsecase := SolutionReactionUsecase.NewDeleteSolutionReactionUsecase(solutionReactionRepository, solutionRepository, tokenGateway)

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

	problemSectorController := controller.NewProblemSectorController(
		tokenGateway,
		createProblemSectorUsecase,
		deleteProblemSectorUsecase,
		getAllSectorsByProblemIdUsecase,
	)

	sectorController := controller.NewSectorController(
		getAllSectorsUsecase,
		getSectorByIdUsecase,
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

	solutionReactionController := controller.NewSolutionReactionController(
		tokenGateway,
		createSolutionReactionUsecase,
		deleteSolutionReactionUsecase,
	)

	authMiddleware := middleware.NewAuthMiddleware(tokenGateway)
	apiKeyMiddleware := middleware.NewApiKeyMiddleware()

	routes.InitializeUserRoutes(&userController, routerGin, authMiddleware, apiKeyMiddleware)
	routes.InitializeProblemsRoutes(&problemController, &problemSectorController, routerGin, authMiddleware, apiKeyMiddleware)
	routes.InitializeSolutionRoutes(&solutionController, &solutionReactionController, routerGin, authMiddleware, apiKeyMiddleware)
	routes.InitializeSectorRoutes(&sectorController, routerGin, authMiddleware, apiKeyMiddleware)

}
