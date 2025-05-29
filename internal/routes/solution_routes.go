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
		routes.GET("/solutions/:problemId", solutionController.GetAllSolutionsByProblemId)
	}

	protected := router.Group("/api")
	protected.Use(apiKeyMiddleware.Apply())
	protected.Use(authMiddleware.Apply())
	{
		protected.POST("/solution", solutionController.CreateSolution)
	}

}
