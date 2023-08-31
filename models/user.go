package models

import (
	"time"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	IsAdmin      bool           `json:"-" gorm:"default:false"`
	Name         string         `json:"name" gorm:"size:128;not null"`
	Username     string         `json:"username" gorm:"size:64;not null;unique"`
	Email        string         `json:"email" gorm:"size:256;not null;unique"`
	Password     []byte         `json:"-" gorm:"not null"`
	Reservations []*Reservation `json:"-" gorm:"many2many:reservation_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
