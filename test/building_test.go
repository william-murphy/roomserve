package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"roomserve/handlers"
	"roomserve/models"
	"testing"
)

func TestCreateBuilding(t *testing.T) {
	// setup
	body := models.NewBuilding{
		Name:    "test building name",
		Address: "test building address",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		t.Fatal("Error encoding building to json")
	}
	req, err := http.NewRequest(http.MethodPost, "/building", &buf)
	if err != nil {
		t.Fatal("Error creating NewRequest")
	}
	rr := httptest.NewRecorder()

	// test
	handlers.CreateBuilding(rr, req)
	if rr.Result().StatusCode != 201 {
		t.Errorf("Status should be 201, got %d", rr.Result().StatusCode)
	}
}
