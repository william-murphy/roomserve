package models

import (
	"time"
)

type Building struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Floors    []Floor   `gorm:"foreignKey:BuildingID" json:"-"`
}
