package handlers

import (
	"net/http"

	"roomserve/database"
	"roomserve/models"
	"roomserve/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateReservation(c *fiber.Ctx) error {
	db := database.DB
	userId := utils.GetUserIdFromCtx(c)
	json := new(models.NewReservation)
	err := c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	newReservation := models.Reservation{
		Title:       json.Title,
		Description: json.Description,
		Start:       json.Start,
		End:         json.End,
		CreatedByID: userId,
		RoomID:      json.RoomID,
	}
	var users []models.User
	if len(json.UserIDs) > 0 {
		db.Find(&users, json.UserIDs)
	}
	err = db.Create(&newReservation).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to create reservation")
	}
	err = db.Model(&newReservation).Association("Users").Replace(users)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid users provided")
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
	id, err := utils.GetIdFromCtx(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid parameter provided")
	}
	reservation := models.Reservation{}
	err = db.Preload("CreatedBy").Preload("Room").Preload("Users").First(&reservation, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Reservation not found")
	}
	return c.Status(http.StatusOK).JSON(reservation)
}

func UpdateReservation(c *fiber.Ctx) error {
	db := database.DB
	id, err := utils.GetIdFromCtx(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid parameter provided")
	}
	json := new(models.NewReservation)
	err = c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}
	reservation := models.Reservation{}
	err = db.First(&reservation, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).SendString("Reservation not found")
	}
	userId := utils.GetUserIdFromCtx(c)
	if reservation.CreatedByID != userId {
		return c.Status(http.StatusUnauthorized).SendString("Not allowed to update this reservation")
	}
	var users []models.User
	if len(json.UserIDs) > 0 {
		db.Find(&users, json.UserIDs)
	}
	err = db.Model(&reservation).Association("Users").Replace(users)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid users provided")
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
