package models

import (
	"time"
)

type Floor struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Name       string    `json:"name"`
	BuildingID uint      `json:"buildingId" gorm:"not null"`
	Rooms      []Room    `json:"-"`
}
