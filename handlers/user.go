package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jacob1225/instagram-bot/database"
	"github.com/jacob1225/instagram-bot/models"
)

func ListUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.DB.Db.Find(&users)

	return c.Status(200).JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}
