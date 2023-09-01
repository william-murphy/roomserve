package router

import (
	"net/http"
	"roomserve/handlers"
	"roomserve/middleware"

	"github.com/go-chi/chi/v5"
)

func Initialize(r *chi.Mux) {

	fs := http.FileServer(http.Dir("./static"))

	r.Use(middleware.Json)

	// serve home page static html
	r.Get("/", func(res http.ResponseWriter, req *http.Request) {
		fs.ServeHTTP(res, req)
	})

	// auth
	r.Post("/auth/login", handlers.LoginUser)
	r.Post("/auth/register", handlers.RegisterUser)

	// admin
	r.Route("/admin", func(r chi.Router) {
		r.Use(middleware.Protected, middleware.Admin)
		r.Get("/", func(res http.ResponseWriter, req *http.Request) {
			fs.ServeHTTP(res, req)
		})
	})

	// building
	r.Route("/building", func(r chi.Router) {
		r.With(middleware.Protected, middleware.Admin).Post("/", handlers.CreateBuilding)
		r.Get("/", handlers.GetBuildings)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handlers.BuildingCtx)
			r.Get("/", handlers.GetBuilding)
			r.With(middleware.Protected, middleware.Admin).Put("/", handlers.UpdateBuilding)
		})
	})

	// floor
	r.Route("/floor", func(r chi.Router) {
		r.With(middleware.Protected, middleware.Admin).Post("/", handlers.CreateFloor)
		r.Get("/", handlers.GetFloors)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handlers.FloorCtx)
			r.Get("/", handlers.GetFloor)
			r.With(middleware.Protected, middleware.Admin).Put("/", handlers.UpdateFloor)
		})
	})

	// room
	r.Route("/room", func(r chi.Router) {
		r.With(middleware.Protected, middleware.Admin).Post("/", handlers.CreateRoom)
		r.Get("/", handlers.GetRooms)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handlers.RoomCtx)
			r.Get("/", handlers.GetRoom)
			r.With(middleware.Protected, middleware.Admin).Put("/", handlers.UpdateRoom)
		})
	})

	/*
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
		userGroup.Get("/reservation", middleware.Protected(), handlers.GetUserReservations)

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
	*/
}
