package database

import (
	"fmt"
	"log"
	"os"

	"roomserve/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error // define error here to prevent overshadowing the global DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to database.")
	}

	err = DB.AutoMigrate(&models.Room{}, &models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
