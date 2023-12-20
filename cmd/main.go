package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jacob1225/instagram-bot/database"
)

func main() {

	database.ConnectDb()
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, From Instagram automator")
    })

    app.Listen(":3000")
}