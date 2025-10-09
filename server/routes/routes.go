package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"server/handlers"
	"server/middlewares"
)

// InitRouter initializes gin router
func InitRouter(log *logrus.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(middlewares.Logger(log), gin.Recovery())
	// router.Use(middlewares.GinContextToContextMiddleware())
	router.Use(middlewares.CORSMiddleware())

	router.GET("/", handlers.RootHandler())
	router.GET("/health", handlers.HealthHandler())
	router.POST("/graphql", handlers.GraphQLHandler())
	router.GET("/playground", handlers.PlaygroundHandler())

	return router
}
