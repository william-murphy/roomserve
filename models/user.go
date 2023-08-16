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
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name" gorm:"not null"`
	Username  string    `json:"username" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Password  []byte    `json:"-" gorm:"not null"`
}
