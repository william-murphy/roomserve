package test

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"roomserve/database"
	"roomserve/handlers"
	"roomserve/models"
	"strings"
	"testing"
)

// CREATE

func TestCreateBuilding(t *testing.T) {
	// setup
	db := database.DB
	tx := db.Begin()
	database.DB = tx

	body := models.NewBuilding{
		Name:    "test building name",
		Address: "test building address",
	}

	// tests
	t.Run("Valid", func(t *testing.T) {
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
		if building.Name != body.Name {
			t.Errorf("New building has incorrect name, is %s but should be %s", building.Name, body.Name)
		}
		if building.Address != body.Address {
			t.Errorf("New building has incorrect address, is %s but should be %s", building.Address, body.Address)
		}
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		body := map[string]interface{}{
			"name": 4,
		}
		req := MockPostRequest(t, body, "/building")
		rr := httptest.NewRecorder()

		// run
		handlers.CreateBuilding(rr, req)

		// test
		if rr.Result().StatusCode != 406 {
			t.Errorf("Status should be 406, got %d", rr.Result().StatusCode)
		}
	})

	t.Run("NameTooLong", func(t *testing.T) {
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
	})

	// teardown
	tx.Rollback()
	database.DB = db
}

// READ

func TestGetBuildings(t *testing.T) {
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
			Address: "jacuzzi",
		},
		{
			Name:    "dolphin",
			Address: "test building address 3",
		},
	}
	err := tx.Create(&newBuildings).Error
	if err != nil {
		t.Fatal("Error creating a building")
	}

	// tests
	t.Run("NoQueryParams", func(t *testing.T) {
		req := MockGetRequest(t, "/building")
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
	})

	t.Run("NameParam", func(t *testing.T) {
		req := MockGetRequest(t, "/building?name=dolphin")
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
		if len(responseStruct) != 1 {
			t.Errorf("Response came back with incorrect data, expected 1 buildings but got %d", len(responseStruct))
		}
		if responseStruct[0].Name != "dolphin" {
			t.Errorf("Received incorrect response, expected dolphin to be dolphin but got %s", responseStruct[0].Name)
		}
	})

	t.Run("AddressParam", func(t *testing.T) {
		req := MockGetRequest(t, "/building?address=jacuzzi")
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
		if len(responseStruct) != 1 {
			t.Errorf("Response came back with incorrect data, expected 1 building but got %d", len(responseStruct))
		}
		if responseStruct[0].Address != "jacuzzi" {
			t.Errorf("Received incorrect response, expected address to be jacuzzi but got %s", responseStruct[0].Address)
		}
	})

	t.Run("LimitParam", func(t *testing.T) {
		req := MockGetRequest(t, "/building?limit=2")
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
		if len(responseStruct) != 2 {
			t.Errorf("Response came back with incorrect data, expected 2 buildings but got %d", len(responseStruct))
		}
	})

	t.Run("InvalidParam", func(t *testing.T) {
		req := MockGetRequest(t, "/building?limit=-1")
		rr := httptest.NewRecorder()

		// run
		handlers.GetBuildings(rr, req)

		// test
		if rr.Result().StatusCode != 400 {
			t.Errorf("Status should be 400, got %d", rr.Result().StatusCode)
		}
	})

	// teardown
	tx.Rollback()
	database.DB = db
}

func TestGetBuilding(t *testing.T) {
	// setup
	db := database.DB
	tx := db.Begin()
	database.DB = tx

	newBuilding := MockBuilding(t)
	ctx := MockContext("building", newBuilding)

	// tests
	t.Run("Valid", func(t *testing.T) {
		req := MockGetRequestWithCtx(t, ctx, "/building", newBuilding.ID)
		rr := httptest.NewRecorder()

		// run
		handlers.GetBuilding(rr, req)

		// test
		if rr.Result().StatusCode != 200 {
			t.Errorf("Status should be 200, got %d", rr.Result().StatusCode)
		}
	})

	t.Run("InvalidCtx", func(t *testing.T) {
		req := MockGetRequestWithCtx(t, context.Background(), "/building", 0)
		rr := httptest.NewRecorder()

		// run
		handlers.GetBuilding(rr, req)

		// test
		if rr.Result().StatusCode != 422 {
			t.Errorf("Status should be 422, got %d", rr.Result().StatusCode)
		}
	})

	// teardown
	tx.Rollback()
	database.DB = db
}

// UPDATE

func TestUpdateBuilding(t *testing.T) {
	// setup
	db := database.DB
	tx := db.Begin()
	database.DB = tx

	newBuilding := MockBuilding(t)
	ctx := MockContext("building", newBuilding)

	body := models.NewBuilding{
		Name:    "test updated building name",
		Address: "test updated building address",
	}

	// tests
	t.Run("Valid", func(t *testing.T) {
		req := MockPutRequest(t, ctx, body, "/building", newBuilding.ID)
		rr := httptest.NewRecorder()

		// run
		handlers.UpdateBuilding(rr, req)

		// test
		if rr.Result().StatusCode != 200 {
			t.Errorf("Status should be 200, got %d", rr.Result().StatusCode)
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
		if building.Name != body.Name {
			t.Errorf("New building has incorrect name, is %s should be %s", building.Name, body.Name)
		}
		if building.Address != body.Address {
			t.Errorf("New building has incorrect address, is %s should be %s", building.Address, body.Address)
		}
	})

	t.Run("InvalidCtx", func(t *testing.T) {
		req := MockPutRequest(t, context.Background(), body, "/building", 0)
		rr := httptest.NewRecorder()

		// run
		handlers.UpdateBuilding(rr, req)

		// test
		if rr.Result().StatusCode != 422 {
			t.Errorf("Status should be 422, got %d", rr.Result().StatusCode)
		}
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		body := map[string]interface{}{
			"name": 4,
		}

		req := MockPutRequest(t, ctx, body, "/building", newBuilding.ID)
		rr := httptest.NewRecorder()

		// run
		handlers.UpdateBuilding(rr, req)

		// test
		if rr.Result().StatusCode != 406 {
			t.Errorf("Status should be 406, got %d", rr.Result().StatusCode)
		}
	})

	t.Run("NameTooLong", func(t *testing.T) {
		body := models.NewBuilding{
			Name:    strings.Repeat("x", 65),
			Address: "test updated building address",
		}

		req := MockPutRequest(t, ctx, body, "/building", newBuilding.ID)
		rr := httptest.NewRecorder()

		// run
		handlers.UpdateBuilding(rr, req)

		// test
		if rr.Result().StatusCode != 400 {
			t.Errorf("Status should be 400, got %d", rr.Result().StatusCode)
		}
	})

	// teardown
	tx.Rollback()
	database.DB = db
}

// DELETE

func TestDeleteBuilding(t *testing.T) {
	// setup
	db := database.DB
	tx := db.Begin()
	database.DB = tx

	newBuilding := MockBuilding(t)
	ctx := MockContext("building", newBuilding)

	// tests
	t.Run("Valid", func(t *testing.T) {
		req := MockDeleteRequest(t, ctx, "/building", newBuilding.ID)
		rr := httptest.NewRecorder()

		// run
		handlers.DeleteBuilding(rr, req)

		// test
		if rr.Result().StatusCode != 204 {
			t.Errorf("Status should be 204, got %d", rr.Result().StatusCode)
		}

		var exists bool
		err := tx.Raw("SELECT EXISTS(SELECT 1 FROM buildings WHERE id = ?)", newBuilding.ID).Scan(&exists).Error
		if err != nil || exists {
			t.Errorf("Building with id %d was not deleted", newBuilding.ID)
		}
	})

	t.Run("InvalidCtx", func(t *testing.T) {
		req := MockDeleteRequest(t, context.Background(), "/building", 0)
		rr := httptest.NewRecorder()

		// run
		handlers.DeleteBuilding(rr, req)

		// test
		if rr.Result().StatusCode != 422 {
			t.Errorf("Status should be 422, got %d", rr.Result().StatusCode)
		}
	})

	// teardown
	tx.Rollback()
	database.DB = db
}

func TestShouldFail(t *testing.T) {
	t.Error("This test should fail, and it did! 1")
}
