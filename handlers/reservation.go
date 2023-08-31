package handlers

// import (
// 	"net/http"

// 	"roomserve/database"
// 	"roomserve/models"
// 	"roomserve/utils"

// 	"gorm.io/gorm"
// )

// func CreateReservation(res http.ResponseWriter, req *http.Request) {
// 	db := database.DB
// 	// get json from request body
// 	json := new(models.NewReservation)
// 	err := c.BodyParser(json)
// 	if err != nil {
// 		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
// 	}

// 	// validate given start and end times
// 	if json.End.Before(json.Start) {
// 		return c.Status(http.StatusBadRequest).SendString("Start time must be before end")
// 	}
// 	if utils.CheckOverlappingTime(json.Start, json.End, json.RoomID) {
// 		return c.Status(http.StatusBadRequest).SendString("Reservation times overlap with existing reservation")
// 	}

// 	// create reservation
// 	userId := utils.GetUserIdFromCtx(c)
// 	newReservation := models.Reservation{
// 		Title:       json.Title,
// 		Description: json.Description,
// 		Start:       json.Start,
// 		End:         json.End,
// 		CreatedByID: userId,
// 		RoomID:      json.RoomID,
// 	}
// 	err = db.Create(&newReservation).Error
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).SendString("Unable to create reservation")
// 	}

// 	// handle users
// 	var users []models.User
// 	if len(json.UserIDs) > 0 {
// 		db.Find(&users, json.UserIDs)
// 	}
// 	err = db.Model(&newReservation).Association("Users").Replace(users)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).SendString("Invalid users provided")
// 	}

// 	return c.Status(http.StatusCreated).JSON(newReservation)
// }

// func GetReservations(res http.ResponseWriter, req *http.Request) {
// 	db := database.DB
// 	Reservations := []models.Reservation{}
// 	db.Model(&models.Reservation{}).Order("ID asc").Limit(100).Find(&Reservations)
// 	return c.Status(http.StatusOK).JSON(Reservations)
// }

// func GetReservation(res http.ResponseWriter, req *http.Request) {
// 	db := database.DB
// 	// validate id param
// 	id, err := utils.GetIdFromCtx(c)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).SendString("Invalid parameter provided")
// 	}

// 	// find reservation with given id in database
// 	reservation := models.Reservation{}
// 	err = db.Preload("CreatedBy").Preload("Room").Preload("Users").First(&reservation, id).Error
// 	if err == gorm.ErrRecordNotFound {
// 		return c.Status(http.StatusNotFound).SendString("Reservation not found")
// 	}
// 	return c.Status(http.StatusOK).JSON(reservation)
// }

// func UpdateReservation(res http.ResponseWriter, req *http.Request) {
// 	db := database.DB
// 	// retrieve and validate reservation id
// 	id, err := utils.GetIdFromCtx(c)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).SendString("Invalid parameter provided")
// 	}

// 	// get reservation from database with given id
// 	reservation := models.Reservation{}
// 	err = db.First(&reservation, id).Error
// 	if err == gorm.ErrRecordNotFound {
// 		return c.Status(http.StatusNotFound).SendString("Reservation not found")
// 	}

// 	// get json from request body
// 	json := new(models.NewReservation)
// 	err = c.BodyParser(json)
// 	if err != nil {
// 		return c.Status(http.StatusNotAcceptable).SendString("Invalid JSON")
// 	}

// 	// check if user id from middleware matches reservation's created by field
// 	userId := utils.GetUserIdFromCtx(c)
// 	if reservation.CreatedByID != userId {
// 		return c.Status(http.StatusUnauthorized).SendString("Not allowed to update this reservation")
// 	}

// 	// validate given start and end times
// 	if json.End.Before(json.Start) {
// 		return c.Status(http.StatusBadRequest).SendString("Start time must be before end")
// 	}
// 	if utils.CheckOverlappingTime(json.Start, json.End, json.RoomID) {
// 		return c.Status(http.StatusBadRequest).SendString("Reservation times overlap with existing reservation")
// 	}

// 	// replace users with given users
// 	var users []models.User
// 	if len(json.UserIDs) > 0 {
// 		db.Find(&users, json.UserIDs)
// 	}
// 	err = db.Model(&reservation).Association("Users").Replace(users)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).SendString("Invalid users provided")
// 	}

// 	// update fields and save
// 	reservation.Title = json.Title
// 	reservation.Description = json.Description
// 	reservation.Start = json.Start
// 	reservation.End = json.End
// 	reservation.RoomID = json.RoomID
// 	err = db.Save(&reservation).Error
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).SendString("Unable to update reservation")
// 	}

// 	return c.Status(http.StatusOK).JSON(reservation)
// }
