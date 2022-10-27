package models

import "gorm.io/gorm"

type Book struct {
	Author string `gorm:"not null"`
	gorm.Model
}
