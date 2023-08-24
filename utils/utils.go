package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetIdFromCtx(c *fiber.Ctx) (uint, error) {
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return 0, errors.New("invalid parameter")
	}
	return uint(id), nil
}
