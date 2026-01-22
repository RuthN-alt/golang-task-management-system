package main

import (
	"github.com/gin-gonic/gin"

	"github.com/RuthN-alt/golang-task-management-system/api/routers"
	"github.com/RuthN-alt/golang-task-management-system/api/initializers"
)

func init() {
	// LoadEnvVariables()
	initializers.LoadEnvVariables()
	// ConnectToDB()
	initializers.ConnectToDB()
	// SyncDatabase()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Golang Task Management System",
		})
	})

	routers.UserRoutes(r)
	routers.TaskRoutes(r)
	r.Run()
}