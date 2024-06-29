package main

import (
	"log"
	"task/database"
	"task/routes"

	_ "task/docs"
)

// @title Time Tracker API
// @version 1.0
// @description API for time tracking
// @host localhost:63342
// @BasePath /
func main() {
	database.InitDatabase()

	r := routes.SetupRouter()
	database.Migrate(database.DB)

	port := "63342"
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
