package utils

import (
	"errors"

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
