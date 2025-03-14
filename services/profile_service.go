package services

import (
	"backend/config"
	"backend/models"
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func UpdateProfile(uid string, request models.UserProfileDetails) (bool, error) {

	client, err := config.FirebaseApp.Firestore(context.Background())
	if err != nil {
		return false, err
	}

	data := map[string]interface{}{
		"name":          request.Name,
		"weight":        request.Weight,
		"height":        request.Height,
		"age":           request.Age,
		"target_weight": request.TargetWeight,
		"activity":      request.Activity,
		"goal":          request.Goal,
		"gender":        request.Gender,
	}

	_, err1 := client.Collection("users").Doc(uid).Set(context.Background(), data, firestore.MergeAll)
	defer client.Close()

	if err1 != nil {
		log.Printf("Error updating user profile: %v", err)
		return false, err1
	}
	return true, nil
}
func GetProfile(uid string) (models.UserProfileDetails, error) {
	cliend, err := config.FirebaseApp.Firestore(context.Background())
	if err != nil {
		return models.UserProfileDetails{}, err
	}
	defer cliend.Close()

	doc, err := cliend.Collection("users").Doc(uid).
		Get(context.Background())
	if err != nil {
		log.Printf("Error getting user profile: %v", err)
		return models.UserProfileDetails{}, err
	}
	var profile models.UserProfileDetails
	log.Printf("Document data: %#v", doc.Data())
	err = doc.DataTo(&profile)
	if err != nil {
		log.Printf("Error parsing user profile: %v", err)
		return models.UserProfileDetails{}, err
	}
	return profile, nil
}
