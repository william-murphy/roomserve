package models

type Room struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Capacity uint   `json:"capacity"`
	FloorID  uint   `json:"floorID"`
}
