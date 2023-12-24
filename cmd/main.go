package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jacob1225/instagram-bot/database"
)

func main() {

	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":80")
}
