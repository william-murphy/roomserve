package handlers

import (
	"net/http"

	"roomserve/database"
	"roomserve/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateFloor(c *fiber.Ctx) error {
	db := database.DB
	json := new(models.Floor)
	if err := c.BodyParser(json); err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	newRoom := models.Floor{
		Name:  json.Name,
		Level: json.Level,
	}
	err := db.Create(&newRoom).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to create floor")
	}
	return c.Status(http.StatusCreated).JSON(newRoom)
}

func GetFloors(c *fiber.Ctx) error {
	db := database.DB
	Rooms := []models.Floor{}
	db.Model(&models.Floor{}).Order("ID asc").Limit(100).Find(&Rooms)
	return c.Status(http.StatusOK).JSON(Rooms)
}

func GetFloor(c *fiber.Ctx) error {
	db := database.DB
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return c.Status(400).SendString("Invalid ID parameter")
	}
	room := models.Floor{}
	query := models.Floor{ID: uint(id)}
	err = db.First(&room, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Floor not found")
	}
	return c.Status(http.StatusOK).JSON(room)
}
