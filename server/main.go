package main

import (
    "log"
    "net/http"

    "server/database"
    "server/handlers"
    "server/models"
    gqlSchema "server/graphql"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/graphql-go/handler"
)

func main() {
    // Initialize database
    database.InitDatabase()
    
    // Run migrations
    if err := models.MigrateUser(database.GetDB()); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
    
    // Initialize Gin router
    r := gin.Default()
    
    // CORS middleware
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"}
    config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
    config.AllowHeaders = []string{"*"}
    r.Use(cors.New(config))
    
    // REST API routes
    api := r.Group("/api")
    {
        api.GET("/users", handlers.GetUsers)
        api.GET("/users/:id", handlers.GetUser)
        api.POST("/users", handlers.CreateUser)
        api.PUT("/users/:id", handlers.UpdateUser)
        api.DELETE("/users/:id", handlers.DeleteUser)
    }
    
    // GraphQL endpoint
    graphqlHandler := handler.New(&handler.Config{
        Schema:   &gqlSchema.Schema,
        Pretty:   true,
        GraphiQL: true,
    })
    
    r.Any("/graphql", gin.WrapH(graphqlHandler))
    
    // Health check endpoint
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":  "ok",
            "message": "Server is running",
        })
    })
    
    // Start server
    log.Println("Server starting on :8080")
    log.Println("REST API: http://localhost:8080/api")
    log.Println("GraphQL Playground: http://localhost:8080/graphql")
    log.Fatal(r.Run(":8080"))
}
