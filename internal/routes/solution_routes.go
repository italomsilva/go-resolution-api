package routes

import (
	"go-resolution-api/internal/controller"
	"go-resolution-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeSolutionRoutes(
	solutionController *controller.SolutionController,
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
	}

}
