package models

type Room struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Capacity uint
}
