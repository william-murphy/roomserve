package models

import (
	"time"
)

type DisplayUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

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
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"size:128;not null"`
	Username  string `gorm:"size:64;not null"`
	Email     string `gorm:"size:256;not null"`
	Password  []byte `gorm:"not null"`
}
