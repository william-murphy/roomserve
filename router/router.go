package router

import (
	"net/http"

	"roomserve/handlers"
	"roomserve/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {
	router.Static("/", "./static/public")

	router.Use(middleware.Json)

	// auth
	authGroup := router.Group("/auth")
	authGroup.Post("/login", handlers.Login)

	// admin
	adminGroup := router.Group("/admin", middleware.Protected(), middleware.Admin)
	adminGroup.Static("/", "./static/admin")

	// user
	userGroup := router.Group("/user")
	userGroup.Post("/", handlers.CreateUser)

	// building
	buildingAdminGroup := adminGroup.Group("/building")
	buildingAdminGroup.Post("/", handlers.CreateBuilding)
	buildingAdminGroup.Put("/:id", handlers.UpdateBuilding)

	buildingGroup := router.Group("/building")
	buildingGroup.Get("/", handlers.GetBuildings)
	buildingGroup.Get("/:id", handlers.GetBuilding)

	// floor
	floorAdminGroup := adminGroup.Group("/floor")
	floorAdminGroup.Post("/", handlers.CreateFloor)
	floorAdminGroup.Put("/:id", handlers.UpdateFloor)

	floorGroup := router.Group("/floor")
	floorGroup.Get("/", handlers.GetFloors)
	floorGroup.Get("/:id", handlers.GetFloor)

	// room
	roomAdminGroup := adminGroup.Group("/room")
	roomAdminGroup.Post("/", handlers.CreateRoom)
	roomAdminGroup.Put("/:id", handlers.UpdateRoom)

	roomGroup := router.Group("/room")
	roomGroup.Get("/", handlers.GetRooms)
	roomGroup.Get("/:id", handlers.GetRoom)

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
