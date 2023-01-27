package models

import "gorm.io/gorm"

type Admin struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	Level    Level  `json:"level" gorm:"default:0"`
	gorm.Model
}

type Level int

const (
	Librarian = iota
	Administrator
)

var LevelStrings = []string{"[Librarian]", "[Administrator]"}

func (l Level) String() string {
	return LevelStrings[l]
}
