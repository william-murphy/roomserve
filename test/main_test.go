package test

import (
	"os"
	"testing"

	"roomserve/config"
	"roomserve/database"
)

func TestMain(m *testing.M) {
	// setup
	config.Initialize()
	database.Connect()

	// run tests
	exitVal := m.Run()

	// teardown

	// exit
	os.Exit(exitVal)
}
