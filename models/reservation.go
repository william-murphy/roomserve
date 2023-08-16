package models

import (
	"time"
)

type Reservation struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Start       *time.Time `json:"start"`
	End         *time.Time `json:"end"`
}
