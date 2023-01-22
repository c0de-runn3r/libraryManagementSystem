package models

import (
	"gorm.io/gorm"
)

type User struct {
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Email            string `json:"email" gorm:"not null;unique"`
	Password         string `json:"password" gorm:"not null"`
	VerificationCode string `json:"verificationCode"`
	EmailVerified    bool   `json:"verified"`
	Role             Role   `json:"role" gorm:"default:0"`
	gorm.Model
}

type UserResponse struct {
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"verified"`
	Role          Role   `json:"role"`
}

type Role int

const (
	Guest = iota
	Admin
)

var RoleLevelStrings = []string{"[Guest]", "[Admin]"}

func (r Role) String() string {
	return RoleLevelStrings[r]
}
