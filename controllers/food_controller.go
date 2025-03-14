package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"

)

func LogMeal(c *gin.Context) {

	jwtToken := c.Request.Header.Get("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Header Authorization"})
		return
	}

	var food models.Food

	if err := c.BindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}
	// Get user ID from token
	uid, err := utils.VerifyJWT(jwtToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return

	}

	//Log meal to the firestore

	logid, err1 := services.LogMeal(food, uid)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Meal logged successfully", "logid": logid})

}

func DescribeMeal(c *gin.Context){

	var request struct{
		Name string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	jwtToken := c.Request.Header.Get("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Header Authorization"})
		return
	}

	
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}
	// Get user ID from token
	_, err := utils.VerifyJWT(jwtToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return

	}


	//Use AI to describe the meal and return the description
	//Returns food model //AI.DescribeMeal(request.Name, request.Description)

	food := services.GeenerateMealDetails(request.Name, request.Description)
	//This is a placeholder for the AI service

	c.JSON(http.StatusOK, gin.H{"message": "Meal described successfully", "food_details":food})//return food model 

}

func SaveMeal(c *gin.Context){
	jwtToken := c.Request.Header.Get("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Header Authorization"})
		return
	}

	var food models.Food

	if err := c.BindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}
	// Get user ID from token
	uid, err := utils.VerifyJWT(jwtToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return

	}

	//Log meal to the firestore

	logid, err1 := services.SaveMeal(food, uid)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Meal Saved successfully", "logid": logid})

}