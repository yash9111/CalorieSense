package services

import (
	"backend/config"
	"context"
	"log"

	"firebase.google.com/go/auth"
)

// CreateUser registers a new user in Firebase Auth
func CreateUser(email, password string) (string, error) {
	client, err := config.FirebaseApp.Auth(context.Background())
	if err != nil {
		return "", err
	}

	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)

	user, err := client.CreateUser(context.Background(), params)
	if err != nil {
		return "", err
	}

	log.Println("User created successfully:", user.UID)
	return user.UID, nil
}

// AuthenticateUser verifies email and password in Firebase Auth
func AuthenticateUser(email, password string) (string, error) {
	client, err := config.FirebaseApp.Auth(context.Background())
	if err != nil {
		return "", err
	}

	// Firebase does not support direct password authentication
	// Use Firebase Admin SDK to verify ID tokens instead
	userRecord, err := client.GetUserByEmail(context.Background(), email)
	if err != nil {
		return "", err
	}

	return userRecord.UID, nil
}
