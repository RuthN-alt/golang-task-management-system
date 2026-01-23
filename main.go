package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/RuthN-alt/golang-task-management-system/api/initializers"
	"github.com/RuthN-alt/golang-task-management-system/api/routers"
)

func init() {
	// Connect to PostgreSQL using Railway-provided environment variables
	initializers.ConnectToDB()

	// Sync database models
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	// Health check endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Golang Task Management System",
		})
	})

	// Register routes
	routers.UserRoutes(r)
	routers.TaskRoutes(r)

	// Listen on PORT from Railway environment
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
