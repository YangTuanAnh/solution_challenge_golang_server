package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "SC2024/routers"
    "SC2024/pkg/firebase"
    "SC2024/pkg/firebase/storage"
    "SC2024/pkg/firebase/firestore"
    "SC2024/pkg/config"
    "context"
)

func main() {
    ctx:= context.Background()
    config.Load()
    firebase.InitializeFirebaseApp()
    firestore.Setup(ctx)
    storage.Setup(ctx)

    app := fiber.New()

    app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format e.g. "localhost, nikschaefer.tech"
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
    routers.Initalize(app)

    app.Listen(":3000")
}