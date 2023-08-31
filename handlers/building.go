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

func BuildingCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		db := database.DB
		// validate building id param
		id, err := strconv.ParseUint(chi.URLParam(req, "id"), 10, 32)
		if err != nil || id < 1 {
			http.Error(res, "Invalid ID parameter", http.StatusBadRequest)
			return
		}

		// get building from database
		var building models.Building
		err = db.Raw("SELECT * FROM buildings WHERE id = ? LIMIT 1", id).Scan(&building).Error
		if err != nil {
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// pass building into request context
		ctx := context.WithValue(req.Context(), "building", building)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

func CreateBuilding(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	// parse json
	reqBody := new(models.NewBuilding)
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusNotAcceptable)
		return
	}

	// create building
	newBuilding := models.Building{
		Name:    reqBody.Name,
		Address: reqBody.Address,
	}
	err = db.Create(&newBuilding).Error
	if err != nil {
		http.Error(res, "Unable to create building", http.StatusBadRequest)
		return
	}

	utils.RespondWithJson(res, 201, newBuilding)
}

func GetBuildings(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	Buildings := []models.Building{}
	db.Model(&models.Building{}).Order("ID asc").Limit(100).Find(&Buildings)
	utils.RespondWithJson(res, 200, Buildings)
}

func GetBuilding(res http.ResponseWriter, req *http.Request) {
	// get building from context and return it as json
	ctx := req.Context()
	building, ok := ctx.Value("building").(*models.Building)
	if !ok {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	utils.RespondWithJson(res, 200, building)
}

func UpdateBuilding(res http.ResponseWriter, req *http.Request) {
	db := database.DB
	// get building from context
	ctx := req.Context()
	building, ok := ctx.Value("building").(*models.Building)
	if !ok {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	// parse json
	reqBody := new(models.NewBuilding)
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusNotAcceptable)
		return
	}

	// update fields and save
	building.Name = reqBody.Name
	building.Address = reqBody.Address
	err = db.Save(&building).Error
	if err != nil {
		http.Error(res, "Could not update building", http.StatusBadRequest)
		return
	}

	utils.RespondWithJson(res, 200, building)
}
