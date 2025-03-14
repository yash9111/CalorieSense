package main

import (
	"backend/config"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"

)

func main() {
	// Initialize Firebase
	config.InitFirebase()

	r := gin.Default()

	routes.AuthRoutes(r)
	routes.ProfileRoutes(r)
	routes.FoodRoutes(r)
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
