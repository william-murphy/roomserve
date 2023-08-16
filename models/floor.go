package models

import (
	"time"
)

type DisplayFloor struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	NumRooms uint            `json:"numRoom"`
	Building DisplayBuilding `json:"building"`
}

type Floor struct {
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string `gorm:"size:64;not null;default:concat('Floor ',currval('floors_id_seq'))"`
	BuildingID uint
	Rooms      []Room
}
