package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"roomserve/database"
	"roomserve/handlers"
	"roomserve/models"
	"strings"
	"testing"
)

// CREATE

func TestCreateBuilding_Valid(t *testing.T) {
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

func TestCreateBuilding_EmptyName(t *testing.T) {
	// setup
	db := database.DB
	tx := db.Begin()
	database.DB = tx

	body := models.NewBuilding{
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

func TestCreateBuilding_NameTooLong(t *testing.T) {
	// setup
	body := models.NewBuilding{
		Name:    strings.Repeat("x", 65),
		Address: "test building address",
	}
	req := MockPostRequest(t, body, "/building")
	rr := httptest.NewRecorder()

	// run
	handlers.CreateBuilding(rr, req)

	// test
	if rr.Result().StatusCode != 400 {
		t.Errorf("Status should be 400, got %d", rr.Result().StatusCode)
	}
}

func TestCreateBuilding_EmptyAddress(t *testing.T) {
	// setup
	db := database.DB
	tx := db.Begin()
	database.DB = tx

	body := models.NewBuilding{
		Name: "test building name",
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

func TestCreateBuilding_AddressTooLong(t *testing.T) {
	// setup
	body := models.NewBuilding{
		Name:    "test building name",
		Address: strings.Repeat("x", 2049),
	}
	req := MockPostRequest(t, body, "/building")
	rr := httptest.NewRecorder()

	// run
	handlers.CreateBuilding(rr, req)

	// test
	if rr.Result().StatusCode != 400 {
		t.Errorf("Status should be 400, got %d", rr.Result().StatusCode)
	}
}

// READ

func TestGetBuildings_NoQueryParams(t *testing.T) {
	// setup
	db := database.DB
	tx := db.Begin()
	database.DB = tx

	newBuildings := []models.Building{
		{
			Name:    "test building name 1",
			Address: "test building address 1",
		},
		{
			Name:    "test building name 2",
			Address: "test building address 2",
		},
		{
			Name:    "test building name 3",
			Address: "test building address 3",
		},
	}
	err := tx.Create(&newBuildings).Error
	if err != nil {
		t.Fatal("Error creating a building")
	}

	req := httptest.NewRequest(http.MethodGet, "/building", nil)
	rr := httptest.NewRecorder()

	// run
	handlers.GetBuildings(rr, req)

	// test
	if rr.Result().StatusCode != 200 {
		t.Errorf("Status should be 200, got %d", rr.Result().StatusCode)
	}

	responseBody := rr.Body.Bytes()
	var responseStruct []models.Building
	if err := json.Unmarshal(responseBody, &responseStruct); err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}
	if len(responseStruct) != 3 {
		t.Errorf("Response came back with incorrect data, expected 3 buildings but got %d", len(responseStruct))
	}

	// teardown
	tx.Rollback()
	database.DB = db
}

func TestGetBuilding_Valid(t *testing.T) {
	// setup
	db := database.DB
	tx := db.Begin()
	database.DB = tx

	newBuilding := models.Building{
		Name:    "test building name",
		Address: "test building address",
	}
	err := tx.Create(&newBuilding).Error
	if err != nil {
		t.Fatal("Error creating a building")
	}

	ctx := context.WithValue(context.Background(), "building", newBuilding)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/building/%d", newBuilding.ID), nil)
	if err != nil {
		t.Fatal("Could not create mock request")
	}
	rr := httptest.NewRecorder()

	// run
	handlers.GetBuilding(rr, req)

	// test
	if rr.Result().StatusCode != 200 {
		t.Errorf("Status should be 200, got %d", rr.Result().StatusCode)
	}

	// teardown
	tx.Rollback()
	database.DB = db
}

// UPDATE

func TestUpdateBuilding_Valid(t *testing.T) {
	// setup

	// run

	// test

	// teardown
}

// DELETE

func TestDeleteBuilding_Valid(t *testing.T) {
	// setup

	// run

	// test

	// teardown
}
