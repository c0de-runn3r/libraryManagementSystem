package models

import (
	"gorm.io/gorm"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email" gorm:"unique"`
	Address string `json:"address"`
	gorm.Model
}
