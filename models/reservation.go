package models

import (
	"time"
)

type CreateReservation struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	CreatedByID uint      `json:"createdById"`
	RoomID      uint      `json:"roomId"`
	UserIDs     []uint    `json:"userIds"`
}

type UpdateReservation struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	RoomID      uint      `json:"roomId"`
	UserIDs     []uint    `json:"userIds"`
}

type Reservation struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Title       string    `json:"title" gorm:"size:64"`
	Description string    `json:"description"`
	Start       time.Time `json:"start" gorm:"not null"`
	End         time.Time `json:"end" gorm:"not null"`
	CreatedByID uint      `json:"-"`
	CreatedBy   *User     `json:"createdBy,omitempty" gorm:"foreignKey:CreatedByID"`
	RoomID      uint      `json:"-"`
	Room        *Room     `json:"room,omitempty" gorm:"foreignKey:RoomID"`
	Users       []*User   `json:"users" gorm:"many2many:reservation_users;"`
}
