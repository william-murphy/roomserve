package models

import (
	"time"
)

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
