package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"

)

// AuthRoutes sets up authentication-related endpoints
func ProfileRoutes(router *gin.Engine) {
	auth := router.Group("/profile")
	{
		auth.GET("/", controllers.GetProfile)
		auth.POST("/", controllers.UpdateProfile)
	}
}
