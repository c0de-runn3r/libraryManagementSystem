// TODO повістю неправильний файл
package models

import "gorm.io/gorm"

type Author struct {
	Name  string  `gorm:"not null"`
	Books []*Book `gorm:"many2many:author_books;"`
	gorm.Model
}
