package models

import (
	"time"
)

type Floor struct {
	ID         uint      `gorm:"primaryKey"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Name       string    `json:"name"`
	Level      uint      `json:"level"`
	BuildingID uint      `json:"buildingId"`
	Rooms      []Room    `gorm:"foreignKey:FloorID" json:"-"`
}
