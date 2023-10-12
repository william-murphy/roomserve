package test

import (
	"encoding/json"
	"net/http/httptest"
	"roomserve/database"
	"roomserve/handlers"
	"roomserve/models"
	"testing"
)

func TestCreateBuilding(t *testing.T) {
	// setup
	db := database.DB
	tx := db.Begin()
	database.DB = tx

	body := models.NewBuilding{
		Name:    "test building name",
		Address: "test building address",
	}
	req := MockPostRequest(t, body, "/building")
	rr := httptest.NewRecorder()

	// run
	handlers.CreateBuilding(rr, req)

	// test
	if rr.Result().StatusCode != 201 {
		t.Errorf("Status should be 201, got %d", rr.Result().StatusCode)
	}

	responseBody := rr.Body.Bytes()
	var responseStruct models.Building
	if err := json.Unmarshal(responseBody, &responseStruct); err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	var building models.Building
	err := tx.Raw("SELECT * FROM buildings WHERE id = ? LIMIT 1", responseStruct.ID).Scan(&building).Error
	if err != nil || building.ID == 0 {
		t.Errorf("New building doesn't exist in database, id: %d", responseStruct.ID)
	}

	// teardown
	tx.Rollback()
	database.DB = db
}
