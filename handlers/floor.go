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
	json := new(models.CreateFloor)
	err := c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	newFloor := models.Floor{
		Name:       json.Name,
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
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID parameter")
	}
	floor := models.Floor{}
	query := models.Floor{ID: uint(id)}
	err = db.First(&floor, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Floor not found")
	}
	return c.Status(http.StatusOK).JSON(floor)
}

func UpdateFloor(c *fiber.Ctx) error {
	db := database.DB
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID parameter")
	}
	json := new(models.CreateFloor)
	err = c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	floor := models.Floor{}
	query := models.Floor{ID: uint(id)}
	err = db.First(&floor, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Floor not found")
	}
	floor.Name = json.Name
	floor.BuildingID = json.BuildingID
	err = db.Save(&floor).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to update floor")
	}
	return c.Status(http.StatusOK).JSON(floor)
}
