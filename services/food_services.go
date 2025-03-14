package services

import (
	"backend/config"
	"backend/models"
	"context"
	"log"
	"time"
)

func LogMeal(foodDetails models.Food, userID string) (string, error) {

	foodDetails.TimeStamp = time.Now()
	client, err := config.FirebaseApp.Firestore(context.Background())
	if err != nil {
		return "", err
	}
	defer client.Close()
	// Reference Firestore collection: /food_logs/{userId}/entries/
	docRef, _, err := client.Collection("food_logs").Doc(userID).Collection("entries").Add(context.Background(), foodDetails)
	if err != nil {
		log.Println("Error logging food entry:", err)
		return "", err
	}

	return docRef.ID, nil

}

func SaveMeal(foodDetails models.Food, userID string) (string, error) {

	foodDetails.TimeStamp = time.Now()
	client, err := config.FirebaseApp.Firestore(context.Background())
	if err != nil {
		return "", err
	}
	defer client.Close()
	// Reference Firestore collection: /food_logs/{userId}/entries/
	docRef, _, err := client.Collection("food_logs").Doc(userID).Collection("saved_meals").Add(context.Background(), foodDetails)
	if err != nil {
		log.Println("Error Saving food entry:", err)
		return "", err
	}

	return docRef.ID, nil

}

func GeenerateMealDetails(name, description string) models.Food {
	//Use AI service to genrate food model based on name and description
	//This is a placeholder for the AI service
	food := models.Food{
		Name:        name,
		Calories:    "200",
		Protein:     "20",
		Carbs:       "30",
		Fat:         "10",
		Fiber:       "5",
		ImageUrl:    "https://www.google.com",
		Ingredients: []string{"rice", "dal"},
		TimeStamp:   time.Now(),
	}
	return food
}
