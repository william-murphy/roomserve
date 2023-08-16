package models

import (
	"time"
)

type Building struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"size:64;not null;default:concat('Building ',currval('buildings_id_seq'))"`
	Floors    []Floor
}
