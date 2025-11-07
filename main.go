package main

import (
	"go-resolution-api/database"
	"go-resolution-api/internal/injection"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseConnection := database.ConnectDB()
	if databaseConnection == nil {
		panic("Error opening connection to the database")
	}

	router := gin.Default()
	routes := router.Group("/api")
	routes.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Log de hello",
		})
	})
	injection.InjectDependencies(databaseConnection, router)
	log.Println("server is running in http://localhost:3060/api")
	router.Run(":3060")
}
