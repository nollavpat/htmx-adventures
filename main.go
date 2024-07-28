package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/nollavpat/chat-htmx/handler"
)

func main() {
	// Create views engine
	viewsEngine := html.New("./views", ".html")

	// Start new fiber instance
	app := fiber.New(fiber.Config{Views: viewsEngine})

	// Static route and directory
	app.Static("/static/", "./static")

	// Create a "ping" handler to test the server
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to fiber")
	})

	appHandler := handler.NewAppHandler()

	app.Get("/", appHandler.HandleGetIndex)

	// create new webscoket
	server := NewWebSocket()
	app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
		server.HandleWebSocket(ctx)
	}))

	go server.HandleMessages()

	// Start the http server
	app.Listen("localhost:3000")
}
