package handlers

import (
	"net/http"
	"roomserve/database"
	"roomserve/models"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	db := database.DB
	json := new(models.LoginUser)
	err := c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
}
