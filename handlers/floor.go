package handlers

import (
	"net/http"

	"roomserve/database"
	"roomserve/models"
	"roomserve/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateFloor(c *fiber.Ctx) error {
	db := database.DB
	// parse json request body
	json := new(models.NewFloor)
	err := c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}

	// create new floor
	newFloor := models.Floor{
		Name:       json.Name,
		Level:      json.Level,
		BuildingID: json.BuildingID,
	}
	err = db.Create(&newFloor).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to create floor")
	}
	return c.Status(http.StatusCreated).JSON(newFloor)
}

func GetFloors(c *fiber.Ctx) error {
	db := database.DB
	Floors := []models.Floor{}
	db.Model(&models.Floor{}).Order("ID asc").Limit(100).Find(&Floors)
	return c.Status(http.StatusOK).JSON(Floors)
}

func GetFloor(c *fiber.Ctx) error {
	db := database.DB
	// validate id param
	id, err := utils.GetIdFromCtx(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid parameter provided")
	}

	// find floor with given id in database
	floor := models.Floor{}
	err = db.Preload("Building").First(&floor, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Floor not found")
	}
	return c.Status(http.StatusOK).JSON(floor)
}

func UpdateFloor(c *fiber.Ctx) error {
	db := database.DB
	// validate id param
	id, err := utils.GetIdFromCtx(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid parameter provided")
	}

	// parse json request body
	json := new(models.NewFloor)
	err = c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}

	// find floor in database
	floor := models.Floor{}
	err = db.First(&floor, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Floor not found")
	}

	// update fields
	floor.Name = json.Name
	floor.Level = json.Level
	floor.BuildingID = json.BuildingID
	err = db.Save(&floor).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to update floor")
	}
	return c.Status(http.StatusOK).JSON(floor)
}
