package routes

import (
	"go-resolution-api/internal/controller"
	"go-resolution-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeProblemsRoutes(
	problemController *controller.ProblemController,
	problemSectorController *controller.ProblemSectorController,
	router *gin.Engine,
	authMiddleware *middleware.AuthMiddleware,
	apiKeyMiddleware *middleware.ApiKeyMiddleware,
) {
	routes := router.Group("/api")
	routes.Use(apiKeyMiddleware.Apply())
	{
		routes.GET("/problems", problemController.GetAllProblems)
		routes.GET("/problem/:problemId", problemController.GetProblemById)
		routes.GET("/problem/:problemId/sectors", problemSectorController.GetSectorsByProblemID)
	}

	protected := router.Group("/api")
	protected.Use(apiKeyMiddleware.Apply())
	protected.Use(authMiddleware.Apply())
	{
		protected.POST("/problem", problemController.CreateProblem)
		protected.PUT("/problem", problemController.UpdateProblem)
		protected.GET("/problems/user", problemController.GetAllProblemsByUserId)
		protected.DELETE("/problem", problemController.DeleteProblem)
		protected.DELETE("/problems/user", problemController.DeleteAllProblemsByUserId)
		protected.POST("/problem/sector", problemSectorController.CreateProblemSector)
		protected.DELETE("/problem/sector", problemSectorController.DeleteProblemSector)
	}

}
