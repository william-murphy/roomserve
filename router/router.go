package router

import (
	"net/http"

	"roomserve/handlers"
	"roomserve/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Hello, World!")
	})

	router.Use(middleware.Json)

	roomGroup := router.Group("/rooms")
	roomGroup.Post("/", handlers.CreateRoom)
	roomGroup.Get("/", handlers.GetRooms)
	roomGroup.Get("/:id", handlers.GetRoom)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	})
}
