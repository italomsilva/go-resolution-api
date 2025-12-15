package routes

import (
	"go-resolution-api/internal/controller"
	"go-resolution-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeSolutionRoutes(
	solutionController *controller.SolutionController,
	solutionReactionController *controller.SolutionReactionController,
	router *gin.Engine,
	authMiddleware *middleware.AuthMiddleware,
	apiKeyMiddleware *middleware.ApiKeyMiddleware,
) {
	routes := router.Group("/api")
	routes.Use(apiKeyMiddleware.Apply())
	{
		routes.GET("/solutions/problem/:problemId", solutionController.GetAllSolutionsByProblemId)
		routes.GET("/solutions/:solutionId", solutionController.GetSolutionById)
	}

	protected := router.Group("/api")
	protected.Use(apiKeyMiddleware.Apply())
	protected.Use(authMiddleware.Apply())
	{
		protected.POST("/solution", solutionController.CreateSolution)
		protected.PUT("/solution", solutionController.UpdateSolution)
		protected.DELETE("/solution", solutionController.DeleteSolution)
		protected.DELETE("/solutions/problem", solutionController.DeleteAllSolutionsByProblemId)
		protected.DELETE("/solutions/user", solutionController.DeleteAllSolutionsByUserId)

		protected.POST("/solution/reaction", solutionReactionController.CreateSolutionReaction)
		protected.DELETE("/solution/reaction", solutionReactionController.DeleteSolutionReaction)
		protected.GET("/solutions/problem/app/:problemId", solutionController.GetAllSolutionsByProblemIdToApp)
		protected.GET("/solutions/user/app", solutionController.GetMySolutionsToApp)

		protected.PUT("/solution/approve", solutionController.ApproveSolution)

		protected.GET("/solution/stats/solutions-reactions", solutionController.StatsCountSolutionsReactions)
	}

}
