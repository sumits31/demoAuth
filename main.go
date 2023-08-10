package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sumits31/demoAuth/authentication/config"
	"github.com/sumits31/demoAuth/authentication/handlers"
	"github.com/sumits31/demoAuth/authentication/middlewares"
)

func main() {
	fmt.Printf("hello fleekians")
	// Create a new Fiber instance
	app := fiber.New()
	// Create a new JWT middleware
	// Note: This is just an example, please use a secure secret key
	jwt := middlewares.NewAuthMiddleware(config.Secret)
	// Create a Login route
	app.Post("/login", handlers.Login)
	// Create a protected route
	app.Get("/protected", jwt, handlers.Protected)
	// Listen on port 3000
	app.Listen(":3000")
}
