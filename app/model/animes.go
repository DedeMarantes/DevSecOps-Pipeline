package model

type Anime struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Title  string `gorm:"not null" json:"title" binding:"required"`
	Genre  string `gorm:"not null" json:"genre"`
	Year   int    `gorm:"not null" json:"year"`
	Author string `json:"author"`
}
