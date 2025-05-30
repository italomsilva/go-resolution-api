package main

import (
	"go-resolution-api/database"
	"go-resolution-api/internal/injection"
	"log"

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
	
	injection.InjectDependencies(databaseConnection, router)

	router.Run(":3060")
}
