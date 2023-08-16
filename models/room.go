package models

import (
	"time"
)

type Room struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Capacity  uint      `json:"capacity" gorm:"not null;check:capacity > 0"`
	FloorID   uint      `json:"floorId" gorm:"not null"`
}
