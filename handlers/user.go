package handlers

import (
	"net/http"

	"roomserve/database"
	"roomserve/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	// parse json request body
	json := new(models.RegisterUser)
	err := c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}

	// create a hash of the given password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to register user")
	}

	// create user (with hashed password)
	newRoom := models.User{
		Name:     json.Name,
		Username: json.Username,
		Email:    json.Email,
		Password: hashedPassword,
	}
	err = db.Create(&newRoom).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to register user")
	}
	return c.Status(http.StatusCreated).JSON(newRoom)
}
