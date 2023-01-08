package models

import (
	"gorm.io/gorm"
)

type Book struct {
	Title     string    `gorm:"not null;unique" json:"title"`
	Authors   []*Author `gorm:"many2many:author_books;" json:"authors"`
	Genres    string    `json:"genres"`
	Publisher string    `json:"publisher"`
	Year      string    `json:"year"`
	ISBN      string    `json:"isbn"`
	UDKNumber string    `json:"udk_number"`
	PageCount int       `json:"page_count"`
	gorm.Model
}
