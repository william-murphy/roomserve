package main

import (
	"log"

	"roomserve/config"
	"roomserve/database"
	"roomserve/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.Connect()
	router.Initialize(app)
	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
