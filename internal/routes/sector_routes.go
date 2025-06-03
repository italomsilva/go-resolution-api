package routes

import (
	"go-resolution-api/internal/controller"
	"go-resolution-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeSectorRoutes(
	sectorController *controller.SectorController,
	router *gin.Engine,
	authMiddleware *middleware.AuthMiddleware,
	apiKeyMiddleware *middleware.ApiKeyMiddleware,
) {
	routes := router.Group("/api")
	routes.Use(apiKeyMiddleware.Apply())
	{
		routes.GET("/sectors", sectorController.GetAllSectors)
		routes.GET("/sector/:sectorId", sectorController.GetSectorById)
	}

}
