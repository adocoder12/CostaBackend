package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CorsMiddleware handles Cross-Origin Resource Sharing.
// Must use a specific origin (not wildcard) when the frontend sends
// credentials (Authorization: Bearer <jwt>).
func CorsMiddleware(allowedOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", allowedOrigin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
