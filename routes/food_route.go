package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(router *gin.Engine) {
	food := router.Group("/food")
	{
		food.POST("/log_meal", controllers.LogMeal)
		food.POST("/describe_meal", controllers.DescribeMeal)
		food.POST("/save_meal", controllers.SaveMeal)

	}
}
