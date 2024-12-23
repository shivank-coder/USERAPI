package main

import (
	"go-practice-app/config"
	"go-practice-app/database"
	"go-practice-app/routes"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to the database
	database.Connect()

	// Set up routes
	r := routes.SetUpRouter()
	// Start the server
	r.Run(":8080")
}
