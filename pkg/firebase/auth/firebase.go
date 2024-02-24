package auth

import (
    "context"
    "github.com/gofiber/fiber/v2"
    "log"
    "net/http"
    "strings"
	"SC2024/pkg/firebase"
)

func FirebaseAuthMiddleware(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return c.Status(http.StatusUnauthorized).SendString("Authorization header is missing")
    }

    tokenParts := strings.Split(authHeader, " ")
    if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
        return c.Status(http.StatusBadRequest).SendString("Invalid Authorization header format")
    }
    jwtToken := tokenParts[1]
	
    firebaseApp := firebase.GetFirebaseApp()

    client, err := firebaseApp.Auth(context.Background())
    if err != nil {
        log.Fatalf("error getting Auth client: %v\n", err)
        return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
    }

    decodedToken, err := client.VerifyIDToken(context.Background(), jwtToken)
    if err != nil {
        log.Printf("error verifying ID token: %v\n", err)
        return c.Status(http.StatusUnauthorized).SendString("Unauthorized")
    }
	c.Locals("user", decodedToken)

    return c.Next()
}