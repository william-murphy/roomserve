package handlers

import (
	"net/http"
	"roomserve/config"
	"roomserve/database"
	"roomserve/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

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
	if !CheckPasswordHash(json.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	// create jwt token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"token": t})
}
