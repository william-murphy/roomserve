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
	json := new(models.NewRoom)
	err := c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	newRoom := models.Room{
		Name:     json.Name,
		Capacity: json.Capacity,
		FloorID:  json.FloorID,
	}
	err = db.Create(&newRoom).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to create room")
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
		return c.Status(http.StatusBadRequest).SendString("Invalid ID parameter")
	}
	room := models.Room{}
	err = db.Preload("Floor").First(&room, uint(id)).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Room not found")
	}
	return c.Status(http.StatusOK).JSON(room)
}

func UpdateRoom(c *fiber.Ctx) error {
	db := database.DB
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID parameter")
	}
	json := new(models.NewRoom)
	err = c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	room := models.Room{}
	err = db.First(&room, uint(id)).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Room not found")
	}
	room.Name = json.Name
	room.Capacity = json.Capacity
	room.FloorID = json.FloorID
	err = db.Save(&room).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to update room")
	}
	return c.Status(http.StatusOK).JSON(room)
}
