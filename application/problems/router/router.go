package router

import (
	"go-resolution-api/application/problems/controller"
	"go-resolution-api/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(problemController *controller.ProblemController, router *gin.Engine) {
	routes := router.Group("/api")
	routes.Use(middleware.ApiKeyMiddleware())
	{
		routes.GET("/problems", problemController.GetAllProblems)
		routes.GET("/problem/:id", problemController.GetProblemById)
	}
	
	protected := router.Group("/api")
	protected.Use(middleware.ApiKeyMiddleware())
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.POST("/problem/", problemController.CreateProblem)
	}

}
