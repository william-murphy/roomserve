package handlers

import (
	"net/http"

	"roomserve/database"
	"roomserve/models"
	"roomserve/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	// parse json request body
	json := new(models.RegisterUser)
	err := c.BodyParser(json)
	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
	}

	// create a hash of the given password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to register user")
	}

	// create user (with hashed password)
	newRoom := models.User{
		Name:     json.Name,
		Username: json.Username,
		Email:    json.Email,
		Password: hashedPassword,
	}
	err = db.Create(&newRoom).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Unable to register user")
	}
	return c.Status(http.StatusCreated).JSON(newRoom)
}

func GetUserReservations(c *fiber.Ctx) error {
	db := database.DB
	// get reservations based on the user id passed via auth
	userId := utils.GetUserIdFromCtx(c)

	// get reservations that include this user
	Reservations := []models.Reservation{}
	db.Raw("SELECT reservations.*,"+
		"rooms.id AS \"Room__id\", rooms.name AS \"Room__name\", rooms.capacity AS \"Room__capacity\", "+
		"floors.id AS \"Room__Floor__id\", floors.name AS \"Room__Floor__name\", floors.level AS \"Room__Floor__level\", "+
		"buildings.id AS \"Room__Floor__Building__id\", buildings.name AS \"Room__Floor__Building__name\" "+
		"FROM reservations LEFT JOIN rooms ON reservations.room_id = rooms.id "+
		"LEFT JOIN floors ON rooms.floor_id = floors.id "+
		"LEFT JOIN buildings ON floors.building_id = buildings.id "+
		"WHERE reservations.id IN (SELECT reservation_id FROM reservation_users WHERE user_id = ?)", userId).Scan(&Reservations)
	return c.Status(http.StatusOK).JSON(Reservations)
}
