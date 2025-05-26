package router

import (
	"go-resolution-api/application/user/controller"
	"go-resolution-api/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(userController *controller.UserController, router *gin.Engine){
	routes := router.Group("/api")
	routes.Use(middleware.ApiKeyMiddleware())
	{
		routes.GET("/users", userController.GetUsers)
		routes.GET("/user/:id", userController.GetUserById)
		routes.POST("/sign-up", userController.CreateUser)
		routes.POST("/sign-in", userController.Login)
	}


	protected := router.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	protected.Use(middleware.ApiKeyMiddleware())
	{
		protected.PUT("/user", userController.UpdateUser)
		protected.DELETE("/user", userController.DeleteAccount)
	}
}