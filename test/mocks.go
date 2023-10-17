package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func MockGetRequest(t *testing.T, route string) *http.Request {
	req := httptest.NewRequest(http.MethodGet, route, nil)
	return req
}

func MockBuildingContext(t *testing.T, newBuilding *models.Building) context.Context {
	db := database.DB
	err := db.Create(newBuilding).Error
	if err != nil {
		t.Fatal("Error creating context")
	}
	ctx := context.WithValue(context.Background(), "building", *newBuilding)
	return ctx
}

func MockGetRequestWithCtx(t *testing.T, ctx context.Context, route string, id uint) *http.Request {
	usableRoute := fmt.Sprintf("%s/%d", route, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, usableRoute, nil)
	if err != nil {
		t.Fatal("Could not create mock get request")
	}
	return req
}

func MockPutRequest(t *testing.T, ctx context.Context, body interface{}, route string, id uint) *http.Request {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		t.Fatal("Error encoding body to json")
	}
	usableRoute := fmt.Sprintf("%s/%d", route, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, usableRoute, &buf)
	if err != nil {
		t.Fatal("Could not create mock put request")
	}
	return req
}

func MockDeleteRequest(t *testing.T, ctx context.Context, route string, id uint) *http.Request {
	usableRoute := fmt.Sprintf("%s/%d", route, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, usableRoute, nil)
	if err != nil {
		t.Fatal("Could not create mock delete request")
	}
	return req
}
