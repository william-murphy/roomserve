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

func RoomCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		db := database.DB
		// validate room id param
		id, err := strconv.ParseUint(chi.URLParam(req, "id"), 10, 32)
		if err != nil || id < 1 {
			http.Error(res, "Invalid ID parameter", http.StatusBadRequest)
			return
		}

		// get room from database
		var room models.Room
		err = db.Raw("SELECT rooms.*, "+
			"floors.id AS \"Floor__id\", floors.name AS \"Floor__name\", floors.level AS \"Floor__level\", "+
			"buildings.id AS \"Floor__Building__id\", buildings.name AS \"Floor__Building__name\", buildings.address AS \"Floor__Building__address\" "+
			"FROM rooms LEFT JOIN floors ON rooms.floor_id = floors.id "+
			"LEFT JOIN buildings ON floors.building_id = buildings.id WHERE rooms.id = ? LIMIT 1", id).Scan(&room).Error
		if err != nil || room.ID == 0 {
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// pass room into request context
		ctx := context.WithValue(req.Context(), "room", room)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

func CreateRoom(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	// parse json
	reqBody := new(models.NewRoom)
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusNotAcceptable)
		return
	}

	// create room
	newRoom := models.Room{
		Name:     reqBody.Name,
		Number:   reqBody.Number,
		Capacity: reqBody.Capacity,
		FloorID:  reqBody.FloorID,
	}
	err = db.Create(&newRoom).Error
	if err != nil {
		http.Error(res, "Unable to create room", http.StatusBadRequest)
		return
	}

	utils.RespondWithJson(res, 201, newRoom)
}

func GetRooms(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	Rooms := []models.Room{}
	db.Raw("SELECT rooms.*, " +
		"floors.id AS \"Floor__id\", floors.name AS \"Floor__name\", floors.level AS \"Floor__level\", " +
		"buildings.id AS \"Floor__Building__id\", buildings.name AS \"Floor__Building__name\", buildings.address AS \"Floor__Building__address\" " +
		"FROM rooms LEFT JOIN floors ON rooms.floor_id = floors.id " +
		"LEFT JOIN buildings ON floors.building_id = buildings.id ORDER BY rooms.id ASC LIMIT 100").Scan(&Rooms)
	utils.RespondWithJson(res, 200, Rooms)
}

func GetRoom(res http.ResponseWriter, req *http.Request) {
	// get room from context and return it as json
	ctx := req.Context()
	room, ok := ctx.Value("room").(models.Room)
	if !ok {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	utils.RespondWithJson(res, 200, room)
}

func UpdateRoom(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	// get room from context
	ctx := req.Context()
	room, ok := ctx.Value("room").(models.Room)
	if !ok {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	// parse json
	reqBody := new(models.NewRoom)
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusNotAcceptable)
		return
	}

	// update fields and save
	room.Name = reqBody.Name
	room.Number = reqBody.Number
	room.Capacity = reqBody.Capacity
	room.FloorID = reqBody.FloorID
	err = db.Save(&room).Error
	if err != nil {
		http.Error(res, "Could not update room", http.StatusBadRequest)
		return
	}

	utils.RespondWithJson(res, 200, room)
}
