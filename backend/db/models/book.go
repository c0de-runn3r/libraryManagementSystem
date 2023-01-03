package models

import (
	"gorm.io/gorm"
)

type Book struct {
	Title     string `gorm:"not null;unique"`
	Authors   string
	Genres    string
	Publisher string
	Year      string
	ISBN      string
	UDKNumber string
	PageCount int
	gorm.Model
}
