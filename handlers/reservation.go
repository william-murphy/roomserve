package handlers

import (
	"net/http"

	"roomserve/database"
	"roomserve/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateReservation(c *fiber.Ctx) error {
	db := database.DB
	json := new(models.Reservation)
	err := c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	newReservation := models.Reservation{
		Title:       json.Title,
		Description: json.Description,
	}
	err = db.Create(&newReservation).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to create room")
	}
	return c.Status(http.StatusCreated).JSON(newReservation)
}

func GetReservations(c *fiber.Ctx) error {
	db := database.DB
	Reservations := []models.Reservation{}
	db.Model(&models.Reservation{}).Order("ID asc").Limit(100).Find(&Reservations)
	return c.Status(http.StatusOK).JSON(Reservations)
}

func GetReservation(c *fiber.Ctx) error {
	db := database.DB
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID parameter")
	}
	room := models.Reservation{}
	query := models.Reservation{ID: uint(id)}
	err = db.First(&room, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Reservation not found")
	}
	return c.Status(http.StatusOK).JSON(room)
}
