package models

import (
	"time"
)

type Building struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name" gorm:"size:64;not null;default:concat('Building ',currval('buildings_id_seq'))"`
	Floors    []Floor   `json:"-"`
}
