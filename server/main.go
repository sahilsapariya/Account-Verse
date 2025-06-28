package main

import (
	"server/config"
	"server/database"
	"server/logs"
	"server/routes"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger := logs.InitLog("info")

	// Initialize database
	err := database.InitDB()
	if err != nil {
		logger.Fatalln("Error initializing the database: ", err)
	}

	// Initialize Gin router with the logger
	r := routes.InitRouter(logger)

	// Start server in a goroutine
	logger.Printf("Server starting on port %s", cfg.Port)
	logger.Printf("GraphQL Playground available at http://localhost:%s/", cfg.Port)
	logger.Fatal(r.Run(":" + cfg.Port))

}
