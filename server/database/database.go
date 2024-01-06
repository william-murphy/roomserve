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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Config("POSTGRES_HOST"), config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DB"), config.Config("POSTGRES_PORT"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to database.")
	}

	err = DB.AutoMigrate(&models.User{}, &models.Building{}, &models.Floor{}, &models.Room{}, &models.Reservation{})
	if err != nil {
		log.Fatal(err)
	}

	err = Seed()
	if err != nil {
		log.Fatal(err)
	}
}
