package routers

import (
	"github.com/gin-gonic/gin"
	
	"github.com/RuthN-alt/golang-task-management-system/api/controllers"
	"github.com/RuthN-alt/golang-task-management-system/api/middleware"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/register", controllers.UserRegistration)
		user.POST("/login", controllers.UserLogin)
		user.PUT("/logout", middleware.AuthMiddleware, controllers.UserLogout)
		user.DELETE("/delete", middleware.AuthMiddleware, controllers.UserDeletion)
	}
}