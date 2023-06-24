package models

type Floor struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Level      uint   `json:"level"`
	BuildingID uint   `json:"buildingId"`
	Rooms      []Room `gorm:"foreignKey:FloorID" json:"-"`
}

// TODO: add gorm column things like max size and not null, etc
