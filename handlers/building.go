package handlers

import (
	"net/http"

	"roomserve/database"
	"roomserve/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateBuilding(c *fiber.Ctx) error {
	db := database.DB
	json := new(models.CreateBuilding)
	err := c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	newRoom := models.Building{
		Name: json.Name,
	}
	err = db.Create(&newRoom).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to create building")
	}
	return c.Status(http.StatusCreated).JSON(newRoom)
}

func GetBuildings(c *fiber.Ctx) error {
	db := database.DB
	Rooms := []models.Building{}
	db.Model(&models.Building{}).Order("ID asc").Limit(100).Find(&Rooms)
	return c.Status(http.StatusOK).JSON(Rooms)
}

func GetBuilding(c *fiber.Ctx) error {
	db := database.DB
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID parameter")
	}
	room := models.Building{}
	query := models.Building{ID: uint(id)}
	err = db.First(&room, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Building not found")
	}
	return c.Status(http.StatusOK).JSON(room)
}
