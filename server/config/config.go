package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Initialize() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Could not load env file")
	}
}

// Config func to get env value
func Config(key string) string {
	return os.Getenv(key)
}
