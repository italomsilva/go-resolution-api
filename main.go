package main

import (
	"go-resolution-api/application/problems"
	"go-resolution-api/application/user"
	"go-resolution-api/database"
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
	problemMain.InitializeModule(databaseConnection, router)
	userMain.InitializeModule(databaseConnection, router)
		
	router.Run(":3060")
}
