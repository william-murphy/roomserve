package models

import (
	"time"
)

type Building struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name" gorm:"unique;not null;default:'Building ' || id"`
	Floors    []Floor   `json:"-"`
}
