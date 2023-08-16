package models

import (
	"time"
)

type Floor struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Name       string    `json:"name" gorm:"size:64;not null;default:concat('Floor ',currval('floors_id_seq'))"`
	BuildingID uint      `json:"buildingId"`
	Rooms      []Room    `json:"-"`
}
