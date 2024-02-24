package controllers

import (
	"github.com/gofiber/fiber/v2"
	"firebase.google.com/go/v4/auth"
	// "context"
	// "SC2024/pkg/firebase"
	// "SC2024/pkg/firebase/auth"
	// "log"
	// "net/http"
)

// Profile screen endpoint
func GetProfile(c *fiber.Ctx) error {
	// Get the JWT token from the request context
	token := c.Locals("user").(*auth.Token)
	// Get the UID of the user from the token
	uid := token.UID
	return c.SendString(uid)
}