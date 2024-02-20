package controllers

import (
	"github.com/gofiber/fiber/v2"
	"SC2024/models"
	"time"
)

// Profile screen endpoint
func GetProfile(c *fiber.Ctx) error {
	// Mock user data
	user := models.User{
		ID: "1",
		Name:  "John Doe",
		Point: 100,
		Activities: []models.Activity{
			{
				ID:    "1",
				Name:  "Activity 1",
				Type:  "streak",
				Date:  time.Now().Format(time.RFC3339),
				Point: 10,
			},
		},
		Achievements: []models.Achievement{
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