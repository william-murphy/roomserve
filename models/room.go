package models

import (
	"time"
)

type NewRoom struct {
	Name     string `json:"name"`
	Number   int16  `json:"number"`
	Capacity uint16 `json:"capacity"`
	FloorID  uint   `json:"floorId"`
}

type Room struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string    `json:"name" gorm:"size:64;default:concat('Room ',currval('rooms_id_seq'))"`
	Number    int16     `json:"number"`
	Capacity  uint16    `json:"capacity" gorm:"not null;default:0"`
	FloorID   uint      `json:"-"`
	Floor     *Floor    `json:"floor"`
}
