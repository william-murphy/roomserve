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
	json := new(models.CreateReservation)
	err := c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	newReservation := models.Reservation{
		Title:       json.Title,
		Description: json.Description,
		Start:       json.Start,
		End:         json.End,
		CreatedByID: json.CreatedByID,
		RoomID:      json.RoomID,
	}
	err = db.Create(&newReservation).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to create reservation")
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
	reservation := models.Reservation{}
	query := models.Reservation{ID: uint(id)}
	err = db.Preload("CreatedBy").Preload("Room").First(&reservation, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Reservation not found")
	}
	return c.Status(http.StatusOK).JSON(reservation)
}

func UpdateReservation(c *fiber.Ctx) error {
	db := database.DB
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID parameter")
	}
	json := new(models.UpdateReservation)
	err = c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	reservation := models.Reservation{}
	query := models.Reservation{ID: uint(id)}
	err = db.First(&reservation, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Reservation not found")
	}
	reservation.Title = json.Title
	reservation.Description = json.Description
	reservation.Start = json.Start
	reservation.End = json.End
	reservation.RoomID = json.RoomID
	err = db.Save(&reservation).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to update reservation")
	}
	return c.Status(http.StatusOK).JSON(reservation)
}

func GetReservationUsers(c *fiber.Ctx) error {
	db := database.DB
	// id, err := c.ParamsInt("id")
	// if err != nil || id < 1 {
	// 	return c.Status(http.StatusBadRequest).SendString("Invalid ID parameter")
	// }
	var users []models.User
	err := db.Model(&models.User{}).Preload("Reservations").Find(&users).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to get users on this reservation")
	}
	return c.Status(http.StatusOK).JSON(users)
}
