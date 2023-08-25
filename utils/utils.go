package utils

import (
	"errors"
	"roomserve/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compares password with hash
func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

// Gets user id from jwt token in ctx locals
func GetUserIdFromCtx(c *fiber.Ctx) uint {
	return uint((c.Locals("user").(*jwt.Token)).Claims.(jwt.MapClaims)["id"].(float64))
}

// Gets id from url
func GetIdFromCtx(c *fiber.Ctx) (uint, error) {
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return 0, errors.New("invalid parameter")
	}
	return uint(id), nil
}

func CheckOverlappingTime(start time.Time, end time.Time, roomId uint) bool {
	db := database.DB
	var Found bool
	db.Raw("SELECT EXISTS(SELECT 1 FROM reservations WHERE reservations.start <= ? AND reservations.end >= ? AND reservations.room_id = ?) AS found", end, start, roomId).Scan(&Found)
	return Found
}
