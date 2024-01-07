package handlers

import (
	"net/http"

	"roomserve/database"
	"roomserve/models"
	"roomserve/utils"
)

func GetUserReservations(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	// get user from context
	ctx := req.Context()
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	// get reservations that include this user
	Reservations := []models.Reservation{}
	db.Raw("SELECT reservations.*,"+
		"rooms.id AS \"Room__id\", rooms.name AS \"Room__name\", rooms.capacity AS \"Room__capacity\", "+
		"floors.id AS \"Room__Floor__id\", floors.name AS \"Room__Floor__name\", floors.level AS \"Room__Floor__level\", "+
		"buildings.id AS \"Room__Floor__Building__id\", buildings.name AS \"Room__Floor__Building__name\" "+
		"FROM reservations LEFT JOIN rooms ON reservations.room_id = rooms.id "+
		"LEFT JOIN floors ON rooms.floor_id = floors.id "+
		"LEFT JOIN buildings ON floors.building_id = buildings.id "+
		"WHERE reservations.id IN (SELECT reservation_id FROM reservation_users WHERE user_id = ?)", user.ID).Scan(&Reservations)

	utils.RespondWithJson(res, 200, Reservations)
}
