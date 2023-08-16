package models

import (
	"time"
)

type DisplayReservation struct {
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Start       time.Time   `json:"start"`
	End         time.Time   `json:"end"`
	CreatedBy   DisplayUser `json:"createdBy"`
	Room        DisplayRoom `json:"room"`
}

type Reservation struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string `gorm:"size:64"`
	Description string
	Start       time.Time `gorm:"not null"`
	End         time.Time `gorm:"not null"`
	CreatedByID uint
	CreatedBy   User
}
