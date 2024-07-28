package main

import "github.com/gofiber/fiber/v2"

func main() {

	// Start new fiber instance
	app := fiber.New()

	// Create a "ping" handler to test the server
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to fiber")
	})

	// Start the http server
	app.Listen("localhost:3000")
}
