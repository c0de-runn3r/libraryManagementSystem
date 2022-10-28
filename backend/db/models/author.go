// TODO повістю неправильний файл
package models

import "gorm.io/gorm"

type Author struct {
	Author string `gorm:"not null"`
	gorm.Model
}
