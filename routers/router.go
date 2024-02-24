package routers

import (
	"github.com/gofiber/fiber/v2"
	"SC2024/controllers"
	"SC2024/pkg/firebase/auth"
)

func Initalize(app *fiber.App) {
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })
	
	app.Get("/", auth.FirebaseAuthMiddleware, controllers.GetProfile)

	app.Get("/coupons", controllers.GetProfile)
	app.Post("/coupons/:id", controllers.GetProfile)

	app.Get("/challeges", controllers.GetProfile)
	app.Get("/challeges/active", controllers.GetProfile)
	app.Post("/challeges/:id", controllers.GetProfile)

	app.Get("/timer", controllers.GetProfile)
	app.Put("/timer", controllers.GetProfile)
	app.Post("/timer", controllers.GetProfile)
	app.Get("/timer/goal", controllers.GetProfile)

	app.Get("/nicotine/:duration", controllers.GetProfile)
	app.Post("/nicotine", controllers.GetProfile)
}
