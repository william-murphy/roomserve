package models

import (
	"time"
)

type DisplayBuilding struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	NumFloors uint   `json:"numFloors"`
}

type Building struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"size:64;not null;default:concat('Building ',currval('buildings_id_seq'))"`
	Floors    []Floor
}
