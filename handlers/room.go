package handlers

import (
	"net/http"

	"roomserve/database"
	"roomserve/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateRoom(c *fiber.Ctx) error {
	db := database.DB
	json := new(models.Room)
	if err := c.BodyParser(json); err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	newRoom := models.Room{
		Name:     json.Name,
		Capacity: json.Capacity,
	}
	err := db.Create(&newRoom).Error
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	return c.Status(http.StatusCreated).JSON(newRoom)
}

func GetRooms(c *fiber.Ctx) error {
	db := database.DB
	Rooms := []models.Room{}
	db.Model(&models.Room{}).Order("ID asc").Limit(100).Find(&Rooms)
	return c.Status(http.StatusOK).JSON(Rooms)
}

func GetRoom(c *fiber.Ctx) error {
	db := database.DB
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return c.Status(400).SendString("Invalid ID parameter")
	}
	room := models.Room{}
	query := models.Room{ID: uint(id)}
	err = db.First(&room, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Room not found")
	}
	return c.Status(http.StatusOK).JSON(room)
}
