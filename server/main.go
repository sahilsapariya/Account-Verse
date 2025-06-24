package main

import (
	// "context"
	// "net/http"
	// "os"
	// "os/signal"
	"server/config"
	"server/logs"
	"server/routes"
	// "syscall"
	// "time"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger := logs.InitLog("info")

	// Initialize Gin router with the logger
	r := routes.InitRouter(logger)

	// Start server in a goroutine
	logger.Printf("Server starting on port %s", cfg.Port)
	logger.Printf("GraphQL Playground available at http://localhost:%s/", cfg.Port)
	logger.Fatal(r.Run(":" + cfg.Port))

}
