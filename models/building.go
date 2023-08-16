package models

import (
	"time"
)

type CreateBuilding struct {
	Name string `json:"name"`
}
type UpdateBuilding struct {
	CreateBuilding
}

type Building struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string    `json:"name" gorm:"size:64;not null;default:concat('Building ',currval('buildings_id_seq'))"`
	Floors    []Floor   `json:"-"`
}
