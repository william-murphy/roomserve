package models

import (
	"time"
)

type Room struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Capacity  uint      `json:"capacity"`
	FloorID   uint      `json:"floorId"`
}
