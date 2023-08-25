package middleware

import (
	"net/http"
	"roomserve/database"
	"roomserve/models"
	"roomserve/utils"

	"github.com/gofiber/fiber/v2"
)

func Admin(c *fiber.Ctx) error {
	db := database.DB
	userId := utils.GetUserIdFromCtx(c)
	var user models.User
	db.First(&user, userId)
	if !user.IsAdmin {
		return c.Status(http.StatusForbidden).SendString("Admin access only")
	}
	return c.Next()
}
