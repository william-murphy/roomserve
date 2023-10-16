package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"roomserve/database"
	"roomserve/models"
	"testing"
)

func MockPostRequest(t *testing.T, body interface{}, route string) *http.Request {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		t.Fatal("Error encoding body to json")
	}
	req := httptest.NewRequest(http.MethodPost, route, &buf)
	return req
}

func MockBuildingContext(t *testing.T, val *models.Building) context.Context {
	db := database.DB
	err := db.Create(&val).Error
	if err != nil {
		t.Fatal("Error creating context")
	}
	ctx := context.WithValue(context.Background(), "building", *val)
	return ctx
}
