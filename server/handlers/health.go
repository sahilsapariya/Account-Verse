package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles health check
func HealthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	}
}
