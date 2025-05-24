package main

import (
	userMain "go-resolution-api/application/user"
	"go-resolution-api/database"
	"log"

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

	userMain.InitializeModule(databaseConnection)
	
}
