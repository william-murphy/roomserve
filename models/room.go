package models

import (
	"time"
)

type Room struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"size:64;not null;default:concat('Room ',currval('rooms_id_seq'))"`
	Capacity  uint   `gorm:"not null;check:capacity > 0"`
	FloorID   uint
}
