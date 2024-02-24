package controllers

import (
	"github.com/gofiber/fiber/v2"
	"firebase.google.com/go/v4/auth"
	"context"
	"fmt"
	"SC2024/pkg/firebase/firestore"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
	// "SC2024/pkg/firebase/auth"
	"SC2024/models"
	"log"
	"net/http"
	// "reflect"
)

// Profile screen endpoint
func GetProfile(c *fiber.Ctx) error {
	// Get the JWT token from the request context
	token := c.Locals("user").(*auth.Token)
	// // Get the UID of the user from the token
	// tokenType := reflect.TypeOf(*token.Claims)

	// // Iterate over the fields of the token type
	// for i := 0; i < tokenType.NumField(); i++ {
	// 	field := tokenType.Field(i)
	// 	fmt.Println("Field Name:", field.Name)
	// }

	uid := token.UID
	userName := token.Claims["email"].(string)
	// Get the username from the user profile
	//username := token.DisplayName

	client := firestore.Client()

	collection := client.Collection("users")
	docRef := collection.Doc(uid)

	snapshot, err := docRef.Get(context.Background())
    if err != nil {
        if status.Code(err) == codes.NotFound {
			fmt.Println("Document does not exist")
		} else {
			log.Fatalf("Error getting document: %v", err)
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}
    }

	if snapshot.Exists() {
        // Document exists
        var user models.User
		if err := snapshot.DataTo(&user); err != nil {
			log.Fatalf("Error converting user document data: %v", err)
		}

		// Return the user data
		fmt.Printf("User: %+v\n", user)
		return c.JSON(fiber.Map{
            "status": fiber.StatusOK,
            "data":   user,
        })
    } else {
        // Document does not exist
        fmt.Println("Document does not exist")
		// User document does not exist, create a new user
        newUser := models.User{
            ID:   uid,
            Name: userName, // You can set default values for other fields if needed
        }

        // Add the new user document to Firestore
        _, err := client.Collection("users").Doc(uid).Set(context.Background(), newUser)
        if err != nil {
            log.Fatalf("Error creating new user: %v", err)
        }

        fmt.Println("New user created")
		return c.JSON(fiber.Map{
            "status": fiber.StatusOK,
            "data":   newUser,
        })
    }
}