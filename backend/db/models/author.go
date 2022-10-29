// TODO повістю неправильний файл
package models

import "gorm.io/gorm"

type Author struct {
	Name string `gorm:"not null"`
	gorm.Model
}
