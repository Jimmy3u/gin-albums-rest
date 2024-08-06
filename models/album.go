package models

type Album struct {
	ID     int     `gorm:"primaryKey" json:"album_id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
