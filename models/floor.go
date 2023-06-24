package models

type Floor struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Level      uint   `json:"level"`
	BuildingID uint   `json:"buildingId"`
	Rooms      []Room `gorm:"foreignKey:FloorID" json:"-"`
}
