package routes

import (
	"go-resolution-api/internal/controller"
	"go-resolution-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeUserRoutes(
	userController *controller.UserController,
	router *gin.Engine,
	authMiddleware *middleware.AuthMiddleware,
	apiKeyMiddleware *middleware.ApiKeyMiddleware,
) {
	routes := router.Group("/api")
	routes.Use(apiKeyMiddleware.Apply())
	{
		routes.GET("/users", userController.GetUsers)
		routes.GET("/user/:userId", userController.GetUserById)
		routes.POST("/sign-up", userController.CreateUser)
		routes.POST("/sign-in", userController.Login)
	}

	protected := router.Group("/api")
	protected.Use(authMiddleware.Apply())
	protected.Use(apiKeyMiddleware.Apply())
	{
		protected.PUT("/user", userController.UpdateUser)
		protected.DELETE("/user", userController.DeleteAccount)
	}
}
