package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jacob1225/instagram-bot/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Get("users", handlers.ListUsers)

	app.Post("user", handlers.CreateUser)

}
