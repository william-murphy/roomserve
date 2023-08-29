package handlers

import (
	"net/http"

	"roomserve/database"
	"roomserve/models"
	"roomserve/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateRoom(c *fiber.Ctx) error {
	db := database.DB
	// parse json request body
	json := new(models.NewRoom)
	err := c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}

	// create room
	newRoom := models.Room{
		Name:     json.Name,
		Number:   json.Number,
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
	// validate id param
	id, err := utils.GetIdFromCtx(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid parameter provided")
	}

	// find room in database with given id
	room := models.Room{}
	err = db.Preload("Floor").First(&room, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Room not found")
	}
	return c.Status(http.StatusOK).JSON(room)
}

func UpdateRoom(c *fiber.Ctx) error {
	db := database.DB
	// validate id param
	id, err := utils.GetIdFromCtx(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid parameter provided")
	}

	// parse json request body
	json := new(models.NewRoom)
	err = c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}

	// find room with given id in database
	room := models.Room{}
	err = db.First(&room, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Room not found")
	}

	// update fields and save
	room.Name = json.Name
	room.Number = json.Number
	room.Capacity = json.Capacity
	room.FloorID = json.FloorID
	err = db.Save(&room).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to update room")
	}
	return c.Status(http.StatusOK).JSON(room)
}
