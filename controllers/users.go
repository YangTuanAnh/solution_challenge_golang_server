package controllers

import (
	"github.com/gofiber/fiber/v3"
)

// Profile screen endpoint
func getProfile(c *fiber.Ctx) error {
	// Mock user data
	user := User{
		Name:  "John Doe",
		Point: 100,
		Activities: []Activity{
			{
				ID:    "1",
				Name:  "Activity 1",
				Type:  "streak",
				Date:  time.Now().Format(time.RFC3339),
				Point: 10,
			},
		},
		Achievements: []Achievement{
			{
				ID:    "1",
				Name:  "Achievement 1",
				Date:  time.Now().Format(time.RFC3339),
			},
		},
	}

	// Respond with user information
	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   user,
	})
}