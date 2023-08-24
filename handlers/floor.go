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
	json := new(models.NewFloor)
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
	id, err := utils.GetIdFromCtx(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid parameter provided")
	}
	floor := models.Floor{}
	err = db.Preload("Building").First(&floor, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Floor not found")
	}
	return c.Status(http.StatusOK).JSON(floor)
}

func UpdateFloor(c *fiber.Ctx) error {
	db := database.DB
	id, err := utils.GetIdFromCtx(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid parameter provided")
	}
	json := new(models.NewFloor)
	err = c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	floor := models.Floor{}
	err = db.First(&floor, id).Error
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
