package models

import "gorm.io/gorm"

type User struct {
	Name     string
	Surname  string
	Email    string `gorm:"not null;unique"`
	Password []byte `gorm:"not null"`
	Roles    string
	gorm.Model
}
