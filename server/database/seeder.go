package database

import (
	"roomserve/config"
	"roomserve/models"

	"golang.org/x/crypto/bcrypt"
)

func seedDefaultAdmin() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.Config("ADMIN_PASSWORD")), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	email := config.Config("ADMIN_EMAIL")
	user := models.User{
		IsAdmin:  true,
		Name:     "Default Admin",
		Username: "admin",
		Email:    email,
		Password: hashedPassword,
	}
	err = DB.Where(models.User{Email: email}).FirstOrCreate(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func Seed() error {
	var err error
	if err = seedDefaultAdmin(); err != nil {
		return err
	}
	return nil
}
