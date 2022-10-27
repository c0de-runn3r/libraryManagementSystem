package models

import "gorm.io/gorm"

type Book struct {
	Title        string `gorm:"not null"`
	Authors      string `gorm:"not null"`
	Publisher    string `gorm:"not null"`
	Year         string `gorm:"not null"`
	ISBN         string `gorm:"not null;unique"`
	Other_cordes string `gorm:"not null"`
	Page_count   string `gorm:"not null"`
	Genres       string `gorm:"not null"`
	gorm.Model
}
