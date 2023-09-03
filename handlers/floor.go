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

func FloorCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		db := database.DB
		// validate floor id param
		id, err := strconv.ParseUint(chi.URLParam(req, "id"), 10, 32)
		if err != nil || id < 1 {
			http.Error(res, "Invalid ID parameter", http.StatusBadRequest)
			return
		}

		// get floor from database
		var floor models.Floor
		err = db.Raw("SELECT floors.*, "+
			"buildings.id AS \"Building__id\", buildings.name AS \"Building__name\", buildings.address AS \"Building__address\" "+
			"FROM floors LEFT JOIN buildings ON floors.building_id = buildings.id WHERE floors.id = ? LIMIT 1", id).Scan(&floor).Error
		if err != nil || floor.ID == 0 {
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// pass floor into request context
		ctx := context.WithValue(req.Context(), "floor", floor)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

func CreateFloor(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	// parse json
	reqBody := new(models.NewFloor)
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusNotAcceptable)
		return
	}

	// create floor
	newFloor := models.Floor{
		Name:       reqBody.Name,
		Level:      reqBody.Level,
		BuildingID: reqBody.BuildingID,
	}
	err = db.Create(&newFloor).Error
	if err != nil {
		http.Error(res, "Unable to create floor", http.StatusBadRequest)
		return
	}

	utils.RespondWithJson(res, 201, newFloor)
}

func GetFloors(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	// build sql based on query string
	query := req.URL.Query()
	params := []interface{}{}
	sql := "SELECT floors.*, buildings.id AS \"Building__id\", buildings.name AS \"Building__name\", buildings.address AS \"Building__address\" " +
		"FROM floors LEFT JOIN buildings ON floors.building_id = buildings.id WHERE floors.id = floors.id "
	if query.Get("name") != "" {
		sql += "AND floors.name ILIKE ? "
		params = append(params, "%"+query.Get("name")+"%")
	}
	if query.Get("level") != "" {
		sql += "AND floors.level = ? "
		params = append(params, query.Get("level"))
	}
	if query.Get("minLevel") != "" {
		sql += "AND floors.level >= ? "
		params = append(params, query.Get("minLevel"))
	}
	if query.Get("maxLevel") != "" {
		sql += "AND floors.level <= ? "
		params = append(params, query.Get("maxLevel"))
	}
	if query.Get("buildingId") != "" {
		sql += "AND floors.building_id = ? "
		params = append(params, query.Get("buildingId"))
	}
	if query.Get("limit") != "" {
		sql += "ORDER BY floors.id ASC LIMIT ?"
		params = append(params, query.Get("limit"))
	} else {
		sql += "ORDER BY floors.id ASC LIMIT 100"
	}

	// run sql
	Floors := []models.Floor{}
	err := db.Raw(sql, params...).Scan(&Floors).Error
	if err != nil {
		http.Error(res, "Could not get floors from database", http.StatusBadRequest)
		return
	}
	utils.RespondWithJson(res, 200, Floors)
}

func GetFloor(res http.ResponseWriter, req *http.Request) {
	// get floor from context and return it as json
	ctx := req.Context()
	floor, ok := ctx.Value("floor").(models.Floor)
	if !ok {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	utils.RespondWithJson(res, 200, floor)
}

func UpdateFloor(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	// get floor from context
	ctx := req.Context()
	floor, ok := ctx.Value("floor").(models.Floor)
	if !ok {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	// parse json
	reqBody := new(models.NewFloor)
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusNotAcceptable)
		return
	}

	// update fields and save
	floor.Name = reqBody.Name
	floor.Level = reqBody.Level
	floor.BuildingID = reqBody.BuildingID
	err = db.Save(&floor).Error
	if err != nil {
		http.Error(res, "Could not update floor", http.StatusBadRequest)
		return
	}

	utils.RespondWithJson(res, 200, floor)
}
