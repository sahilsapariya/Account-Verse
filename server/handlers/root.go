package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


// RootHandler handles root endpoint
func RootHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/playground")
    }
}