package models

type Room struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Capacity uint   `json:"capacity"`
	FloorID  uint   `json:"floorId"`
}
