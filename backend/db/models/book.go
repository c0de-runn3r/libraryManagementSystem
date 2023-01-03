package models

import (
	"gorm.io/gorm"
)

type Book struct {
	Title     string    `gorm:"not null;unique"`
	Authors   []*Author `gorm:"many2many:author_books;"`
	Genres    string
	Publisher string
	Year      string
	ISBN      string
	UDKNumber string
	PageCount int
	gorm.Model
}
