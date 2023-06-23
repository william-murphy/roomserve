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

	buildingGroup := router.Group("/buildings")
	buildingGroup.Post("/", handlers.CreateBuilding)
	buildingGroup.Get("/", handlers.GetBuildings)
	buildingGroup.Get("/:id", handlers.GetBuilding)

	floorGroup := router.Group("/floors")
	floorGroup.Post("/", handlers.CreateFloor)
	floorGroup.Get("/", handlers.GetFloors)
	floorGroup.Get("/:id", handlers.GetFloor)

	roomGroup := router.Group("/rooms")
	roomGroup.Post("/", handlers.CreateRoom)
	roomGroup.Get("/", handlers.GetRooms)
	roomGroup.Get("/:id", handlers.GetRoom)

	userGroup := router.Group("/users")
	userGroup.Post("/", handlers.CreateUser)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	})
}
