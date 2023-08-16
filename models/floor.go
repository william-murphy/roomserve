package models

import (
	"time"
)

type Floor struct {
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string `gorm:"size:64;not null;default:concat('Floor ',currval('floors_id_seq'))"`
	BuildingID uint
	Rooms      []Room
}
