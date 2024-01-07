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

	// user
	r.Route("/user", func(r chi.Router) {
		r.Use(middleware.Protected)
		r.Get("/reservation", handlers.GetUserReservations)
	})

	// building
	r.Route("/building", func(r chi.Router) {
		r.With(middleware.Protected, middleware.Admin).Post("/", handlers.CreateBuilding)
		r.Get("/", handlers.GetBuildings)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handlers.BuildingCtx)
			r.Get("/", handlers.GetBuilding)
			r.With(middleware.Protected, middleware.Admin).Put("/", handlers.UpdateBuilding)
			r.With(middleware.Protected, middleware.Admin).Delete("/", handlers.DeleteBuilding)
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
			r.With(middleware.Protected, middleware.Admin).Delete("/", handlers.DeleteFloor)
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
			r.With(middleware.Protected, middleware.Admin).Delete("/", handlers.DeleteRoom)
		})
	})

	// reservation
	r.Route("/reservation", func(r chi.Router) {
		r.With(middleware.Protected).Post("/", handlers.CreateReservation)
		r.Get("/", handlers.GetReservations)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handlers.ReservationCtx)
			r.Get("/", handlers.GetReservation)
			r.Get("/user", handlers.GetReservationUsers)
			r.With(middleware.Protected).Put("/", handlers.UpdateReservation)
			r.With(middleware.Protected).Delete("/", handlers.DeleteReservation)
		})
	})

}
