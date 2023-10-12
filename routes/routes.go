package routes

import (
	"ServiceSentinel/handlers" // Import your route handlers
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes and configures the API routes.
func SetupRoutes(router *gin.Engine) {
	// Example route for getting albums
	router.GET("/services", handlers.GetServices)
	router.GET("/services/:url", handlers.HealthHandler)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	// Add more routes here as needed
}
