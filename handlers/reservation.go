package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"roomserve/database"
	"roomserve/models"
	"roomserve/utils"

	"github.com/go-chi/chi/v5"
)

func ReservationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		db := database.DB
		// validate reservation id param
		id, err := strconv.ParseUint(chi.URLParam(req, "id"), 10, 32)
		if err != nil || id < 1 {
			http.Error(res, "Invalid ID parameter", http.StatusBadRequest)
			return
		}

		// get reservation from database
		var reservation models.Reservation
		err = db.Raw("SELECT reservations.*, "+
			"rooms.id AS \"Room__id\", rooms.name AS \"Room__name\", rooms.number AS \"Room__number\", rooms.capacity AS \"Room__capacity\", "+
			"floors.id AS \"Room__Floor__id\", floors.name AS \"Room__Floor__name\", floors.level AS \"Room__Floor__level\", "+
			"buildings.id AS \"Room__Floor__Building__id\", buildings.name AS \"Room__Floor__Building__name\", buildings.address AS \"Room__Floor__Building__address\" "+
			"FROM reservations LEFT JOIN rooms ON reservations.room_id = rooms.id "+
			"LEFT JOIN floors ON rooms.floor_id = floors.id "+
			"LEFT JOIN buildings ON floors.building_id = buildings.id WHERE reservations.id = ? LIMIT 1", id).Scan(&reservation).Error
		if err != nil || reservation.ID == 0 {
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// pass reservation into request context
		ctx := context.WithValue(req.Context(), "reservation", reservation)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

func CreateReservation(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	// parse json
	reqBody := new(models.NewReservation)
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusNotAcceptable)
		return
	}

	// validate given start and end times
	if reqBody.End.Before(reqBody.Start) {
		http.Error(res, "Start time must be before end time", http.StatusBadRequest)
		return
	}
	if utils.CheckOverlappingTime(0, reqBody.Start, reqBody.End, reqBody.RoomID) {
		http.Error(res, "Reservation time overlaps with an existing reservation for the given reservation", http.StatusBadRequest)
		return
	}

	// get user from ctx
	user := req.Context().Value("user").(models.User)

	// handle users
	var users []*models.User
	if len(reqBody.UserIDs) > 0 {
		db.Find(&users, reqBody.UserIDs)
	}
	if utils.ExceedsRoomCapacity(len(users), reqBody.RoomID) {
		http.Error(res, "Number of users exceeds room capacity", http.StatusBadRequest)
		return
	}

	// create reservation
	newReservation := models.Reservation{
		Title:       reqBody.Title,
		Description: reqBody.Description,
		Start:       reqBody.Start,
		End:         reqBody.End,
		CreatedByID: user.ID,
		RoomID:      reqBody.RoomID,
		Users:       users,
	}
	err = db.Create(&newReservation).Error
	if err != nil {
		http.Error(res, "Unable to create reservation", http.StatusBadRequest)
		return
	}

	utils.RespondWithJson(res, 201, newReservation)
}

func GetReservations(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	Reservations := []models.Reservation{}
	db.Raw("SELECT reservations.*, " +
		"rooms.id AS \"Room__id\", rooms.name AS \"Room__name\", rooms.number AS \"Room__number\", rooms.capacity AS \"Room__capacity\", " +
		"floors.id AS \"Room__Floor__id\", floors.name AS \"Room__Floor__name\", floors.level AS \"Room__Floor__level\", " +
		"buildings.id AS \"Room__Floor__Building__id\", buildings.name AS \"Room__Floor__Building__name\", buildings.address AS \"Room__Floor__Building__address\" " +
		"FROM reservations LEFT JOIN rooms ON reservations.room_id = rooms.id " +
		"LEFT JOIN floors ON rooms.floor_id = floors.id " +
		"LEFT JOIN buildings ON floors.building_id = buildings.id ORDER BY reservations.id ASC LIMIT 100").Scan(&Reservations)
	utils.RespondWithJson(res, 200, Reservations)
}

func GetReservation(res http.ResponseWriter, req *http.Request) {
	// get reservation from context and return it as json
	ctx := req.Context()
	reservation, ok := ctx.Value("reservation").(models.Reservation)
	if !ok {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	utils.RespondWithJson(res, 200, reservation)
}

func UpdateReservation(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	// get reservation from context
	ctx := req.Context()
	reservation, ok := ctx.Value("reservation").(models.Reservation)
	if !ok {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	// parse json
	reqBody := new(models.NewReservation)
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusNotAcceptable)
		return
	}

	// get user from context
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	// check if user id from middleware matches reservation's created by field
	if reservation.CreatedByID != user.ID {
		http.Error(res, "Must be reservation creator to update", http.StatusBadRequest)
		return
	}

	// validate given start and end times
	if reqBody.End.Before(reqBody.Start) {
		http.Error(res, "Start time must be before end time", http.StatusBadRequest)
		return
	}
	if utils.CheckOverlappingTime(reservation.ID, reqBody.Start, reqBody.End, reqBody.RoomID) {
		http.Error(res, "Reservation time overlaps with an existing reservation for the given reservation", http.StatusBadRequest)
		return
	}

	// handle users
	var users []*models.User
	if len(reqBody.UserIDs) > 0 {
		db.Find(&users, reqBody.UserIDs)
	}
	if utils.ExceedsRoomCapacity(len(users), reqBody.RoomID) {
		http.Error(res, "Number of users exceeds room capacity", http.StatusBadRequest)
		return
	}
	err = db.Model(&reservation).Association("Users").Replace(users)
	if err != nil {
		http.Error(res, "Invalid users provided", http.StatusBadRequest)
		return
	}

	// update fields and save
	reservation.Title = reqBody.Title
	reservation.Description = reqBody.Description
	reservation.Start = reqBody.Start
	reservation.End = reqBody.End
	reservation.RoomID = reqBody.RoomID
	err = db.Save(&reservation).Error
	if err != nil {
		http.Error(res, "Unable to update reservation", http.StatusBadRequest)
		return
	}

	utils.RespondWithJson(res, 200, reservation)
}
