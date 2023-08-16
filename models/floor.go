package models

import (
	"time"
)

type CreateFloor struct {
	Name       string `json:"name"`
	BuildingID uint   `json:"buildingId"`
}

type UpdateFloor struct {
	CreateFloor
}

type Floor struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	Name       string    `json:"name" gorm:"size:64;not null;default:concat('Floor ',currval('floors_id_seq'))"`
	BuildingID uint      `json:"-"`
	Rooms      []Room    `json:"-"`
}
