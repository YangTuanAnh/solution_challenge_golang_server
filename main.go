package main

import (
    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "SC203"
)

func main() {
    app := fiber.New()

    app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format e.g. "localhost, nikschaefer.tech"
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
    router.Initalize(app)

    app.Listen(":3000")
}