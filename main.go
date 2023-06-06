package main

import (
	"log"
	"os"

	"roomserve/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New()
	database.Connect()
	log.Fatal(app.Listen(":" + os.Getenv("BACKEND_PORT")))
}
