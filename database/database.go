package database

import (
	"fmt"
	"log"

	"roomserve/config"
	"roomserve/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error // define error here to prevent overshadowing the global DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), config.Config("DB_PORT"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to database.")
	}

	err = DB.AutoMigrate(&models.User{}, &models.Building{}, &models.Floor{}, &models.Room{}, &models.Reservation{})
	if err != nil {
		log.Fatal(err)
	}
}
