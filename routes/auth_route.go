package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"

)

// AuthRoutes sets up authentication-related endpoints
func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", controllers.Signup)
		auth.POST("/login", controllers.Login)
	}
}
