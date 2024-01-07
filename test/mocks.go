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
	"time"

	"golang.org/x/crypto/bcrypt"
)

var counter uint = 0
var baseTime time.Time = time.Now()

// Mocks for HTTP

func MockContext(key string, val interface{}) context.Context {
	ctx := context.WithValue(context.Background(), key, val)
	return ctx
}

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

// Mocks for Database Models

func MockBuilding(t *testing.T) models.Building {
	db := database.DB
	newBuilding := models.Building{
		Name:    "test building name",
		Address: "test building address",
	}
	err := db.Create(&newBuilding).Error
	if err != nil {
		t.Fatal("Error creating mock building")
	}
	return newBuilding
}

func MockFloor(t *testing.T) models.Floor {
	db := database.DB
	building := MockBuilding(t)
	newFloor := models.Floor{
		Name:       "test floor name",
		Level:      1,
		BuildingID: building.ID,
	}
	err := db.Create(&newFloor).Error
	if err != nil {
		t.Fatal("Error creating mock floor")
	}
	return newFloor
}

func MockRoom(t *testing.T) models.Room {
	db := database.DB
	floor := MockFloor(t)
	newRoom := models.Room{
		Name:     "test room name",
		Number:   1,
		Capacity: 4,
		FloorID:  floor.ID,
	}
	err := db.Create(&newRoom).Error
	if err != nil {
		t.Fatal("Error creating mock room")
	}
	return newRoom
}

func MockUser(t *testing.T) models.User {
	db := database.DB
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("hunter2"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal("Error hashing password for mock user")
	}
	newUser := models.User{
		Name:     "John Smith",
		Username: fmt.Sprintf("jsmith%d", counter),
		Email:    fmt.Sprintf("john.smith%d@roomserve.com", counter),
		Password: hashedPassword,
	}
	counter++
	err = db.Create(&newUser).Error
	if err != nil {
		t.Fatal("Error creating mock user")
	}
	return newUser
}

func MockReservation(t *testing.T) models.Reservation {
	db := database.DB
	room := MockRoom(t)
	createdBy := MockUser(t)
	user1 := MockUser(t)
	user2 := MockUser(t)
	user3 := MockUser(t)
	users := []*models.User{&user1, &user2, &user3}
	newReservation := models.Reservation{
		Title:       "test reservation title",
		Description: "test reservation description",
		Start:       baseTime.Add(time.Hour * time.Duration(counter)),
		End:         baseTime.Add(time.Hour * time.Duration(counter+1)),
		CreatedByID: createdBy.ID,
		RoomID:      room.ID,
		Users:       users,
	}
	counter++
	err := db.Create(&newReservation).Error
	if err != nil {
		t.Fatal("Error creating mock reservation")
	}
	return newReservation
}
