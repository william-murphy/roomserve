package models

type Floor struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Level uint   `json:"level"`
}
