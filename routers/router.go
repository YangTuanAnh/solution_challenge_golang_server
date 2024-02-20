package routers

import (
	"github.com/gofiber/fiber/v2"
	"SC2024/controllers"
)

func Initalize(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	
	app.Get("/user", controllers.GetProfile)
}
