// TODO Other_cordes заміни на ті коди, які вам тре. І не cordes, а codes.
package models

import (
	"gorm.io/gorm"
)

type Book struct {
	Title      string `gorm:"not null;unique"`
	Authors    string
	Genres     string
	Publisher  string
	Year       string
	ISBN       string
	UDK_number string
	Page_count int
	gorm.Model
}
