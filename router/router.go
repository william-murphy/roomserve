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

	// auth
	authGroup := router.Group("/auth")
	authGroup.Post("/login", handlers.Login)

	// user
	userGroup := router.Group("/user")
	userGroup.Post("/", handlers.CreateUser)

	// building
	buildingGroup := router.Group("/building")
	buildingGroup.Post("/", handlers.CreateBuilding)
	buildingGroup.Get("/", handlers.GetBuildings)
	buildingGroup.Get("/:id", handlers.GetBuilding)
	buildingGroup.Put("/:id", handlers.UpdateBuilding)

	// floor
	floorGroup := router.Group("/floor")
	floorGroup.Post("/", handlers.CreateFloor)
	floorGroup.Get("/", handlers.GetFloors)
	floorGroup.Get("/:id", handlers.GetFloor)
	floorGroup.Put("/:id", handlers.UpdateFloor)

	// room
	roomGroup := router.Group("/room")
	roomGroup.Post("/", handlers.CreateRoom)
	roomGroup.Get("/", handlers.GetRooms)
	roomGroup.Get("/:id", handlers.GetRoom)
	roomGroup.Put("/:id", handlers.UpdateRoom)

	// reservation
	reservationGroup := router.Group("/reservation")
	reservationGroup.Post("/", middleware.Protected(), handlers.CreateReservation)
	reservationGroup.Get("/", handlers.GetReservations)
	reservationGroup.Get("/:id", handlers.GetReservation)
	reservationGroup.Put("/:id", middleware.Protected(), handlers.UpdateReservation)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	})
}
