package models

import "gorm.io/gorm"

type BookLog struct {
	BookID     uint
	Book       Book `gorm:"foreignKey:BookID;references:ID"`
	UserID     uint
	User       User `gorm:"foreignKey:UserID;references:ID"`
	Returned   bool
	gorm.Model // created is a timestamp for renting book and updated is for returning book
}
