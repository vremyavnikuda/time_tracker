package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time_tracker/config"
	"time_tracker/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.InitializeDatabase()
	router := routers.SetupRouter()

	log.Println("Starting server on port", os.Getenv("PORT"))
	router.Run(":" + os.Getenv("PORT"))
}
