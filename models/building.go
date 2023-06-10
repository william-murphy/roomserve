package models

type Building struct {
	ID     uint    `gorm:"primaryKey"`
	Name   string  `json:"name"`
	Floors []Floor `gorm:"foreignKey:BuildingID"`
}
