package main

import (
	"log"
	"weather-check-app/db"
	"weather-check-app/router"
)

func main() {
	// Initialize the database
	db.InitDB("weather.sqlite")

	// Start the server
	r := router.SetupRouter()
	log.Println("Starting server on port 8081...")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
