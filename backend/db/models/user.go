package models

import "gorm.io/gorm"

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password []byte `gorm:"not null" json:"password"`
	Roles    string
	gorm.Model
}
