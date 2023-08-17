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
	newBuilding := models.Building{
		Name: json.Name,
	}
	err = db.Create(&newBuilding).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to create building")
	}
	return c.Status(http.StatusCreated).JSON(newBuilding)
}

func GetBuildings(c *fiber.Ctx) error {
	db := database.DB
	Buildings := []models.Building{}
	db.Model(&models.Building{}).Order("ID asc").Limit(100).Find(&Buildings)
	return c.Status(http.StatusOK).JSON(Buildings)
}

func GetBuilding(c *fiber.Ctx) error {
	db := database.DB
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID parameter")
	}
	building := models.Building{}
	err = db.First(&building, uint(id)).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Building not found")
	}
	return c.Status(http.StatusOK).JSON(building)
}

func UpdateBuilding(c *fiber.Ctx) error {
	db := database.DB
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID parameter")
	}
	json := new(models.CreateBuilding)
	err = c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	building := models.Building{}
	err = db.First(&building, uint(id)).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Building not found")
	}
	building.Name = json.Name
	err = db.Save(&building).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to update building")
	}
	return c.Status(http.StatusOK).JSON(building)
}
