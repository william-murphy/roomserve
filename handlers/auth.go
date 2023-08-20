package handlers

import (
	"net/http"
	"roomserve/config"
	"roomserve/database"
	"roomserve/models"
	"roomserve/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	// parse json
	db := database.DB
	json := new(models.LoginUser)
	err := c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}

	// find user by email
	user := models.User{}
	query := models.User{Email: json.Email}
	err = db.First(&user, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusBadRequest).SendString("Email not found")
	}

	// check password against database
	if !utils.CheckPasswordHash(json.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	// create jwt token
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"token": t})
}
