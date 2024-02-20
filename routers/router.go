package router

import (
	"github.com/gofiber/fiber/v3"
	"controllers"
)

app.Get("/", func(c fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
})

app.Get("/profile", controllers.getProfile)