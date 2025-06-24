package routes

import (
	"server/graph"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"server/config"
	"server/database"
	"server/handlers"
	"server/middlewares"
)

// InitRouter initializes gin router
func InitRouter(log *logrus.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	cfg := config.LoadConfig()
	db := database.NewDatabase(cfg)
	defer db.Close()

	// Initialize GraphQL resolver
	resolver := graph.NewResolver(db)

	router.Use(middlewares.Logger(log), gin.Recovery())
	// router.Use(middlewares.GinContextToContextMiddleware())
	router.Use(middlewares.CORSMiddleware())

	router.GET("/", handlers.RootHandler())
	router.GET("/health", handlers.HealthHandler())
	router.POST("/query", handlers.GraphQLHandler(resolver))
	router.GET("/playground", handlers.PlaygroundHandler())

	return router
}
