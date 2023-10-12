package main

import (
	"ServiceSentinel/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Initialize and configure the API routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
