package handlers

import (
	"net/http"

	"roomserve/database"
	"roomserve/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	json := new(models.User)
	if err := c.BodyParser(json); err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	newRoom := models.User{
		Username: json.Username,
		Email:    json.Email,
		Password: json.Password,
	}
	err := db.Create(&newRoom).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to register user")
	}
	return c.Status(http.StatusCreated).JSON(newRoom)
}
