package models

import "gorm.io/gorm"

type User struct {
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Roles    string
	gorm.Model
}
