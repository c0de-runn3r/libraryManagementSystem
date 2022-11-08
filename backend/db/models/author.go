// TODO повістю неправильний файл
package models

import "gorm.io/gorm"

type Author struct {
	Name []Book `gorm:"many2many:author_books;"`
	gorm.Model
}
