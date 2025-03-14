package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"

)

func GetProfile(c *gin.Context) {
	jwtToken := c.Request.Header.Get("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Header Authorization"})
		return
	}

	// Get user ID from token
	uid, err := utils.VerifyJWT(jwtToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	
	
	}

	// Get user profile
	profile, err := services.GetProfile(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)	

}
func UpdateProfile(c *gin.Context) {

	var request models.UserProfileDetails

	jwtToken := c.Request.Header.Get("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Header Authorization"})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Get user ID from token
	uid, err := utils.VerifyJWT(jwtToken)

	if err !=nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	// Update user profile
	isCreated, err :=services.UpdateProfile(uid, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !isCreated {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
