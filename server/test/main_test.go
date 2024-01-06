package test

import (
	"os"
	"testing"

	"roomserve/config"
	"roomserve/database"
	"roomserve/models"
)

func TestMain(m *testing.M) {
	// setup
	config.Initialize()
	database.Connect()

	// run tests
	exitVal := m.Run()

	// teardown
	database.DB.Migrator().DropTable(&models.User{}, &models.Building{}, &models.Floor{}, &models.Room{}, &models.Reservation{}, "reservation_users")

	// exit
	os.Exit(exitVal)
}
