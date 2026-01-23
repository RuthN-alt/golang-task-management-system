package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/RuthN-alt/golang-task-management-system/api/initializers"
	"github.com/RuthN-alt/golang-task-management-system/api/routers"
)

func init() {
	// Load env variables from Railway or .env
	initializers.LoadEnvVariables()

	// Connect to PostgreSQL
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

	// Listen on PORT from environment (Railway sets this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)
}
