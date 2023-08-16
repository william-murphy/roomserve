package models

import (
	"time"
)

type CreateRoom struct {
	Name     string `json:"name"`
	Capacity uint   `json:"capacity"`
	FloorID  uint   `json:"floorId"`
}

type UpdateRoom struct {
	CreateRoom
}

type Room struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string    `json:"name" gorm:"size:64;not null;default:concat('Room ',currval('rooms_id_seq'))"`
	Capacity  uint      `json:"capacity" gorm:"not null;check:capacity > 0"`
	FloorID   uint      `json:"-"`
}
