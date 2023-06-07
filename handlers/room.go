package handlers

import (
	"roomserve/database"
	"roomserve/models"

	"github.com/gofiber/fiber/v2"
)

func CreateRoom(c *fiber.Ctx) error {
	db := database.DB
	json := new(models.Room)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	newRoom := models.Room{
		Name:     json.Name,
		Capacity: json.Capacity,
	}
	err := db.Create(&newRoom).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}
