package models

type Building struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Name   string  `json:"name"`
	Floors []Floor `gorm:"foreignKey:BuildingID" json:"-"`
}
