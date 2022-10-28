package db

import (
	. "lms/db/models"
	. "lms/utils"

	"os"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	Log("info", "Auto migration starts...")

	err_user := db.AutoMigrate(&User{})
	err_book := db.AutoMigrate(&Book{})
	err_author := db.AutoMigrate(&Author{})

	if err_user != nil {
		Log("error", "User migration error.\n"+err_user.Error())
		os.Exit(1)
	}
	if err_book != nil {
		Log("error", "Book migration error.\n"+err_book.Error())
		os.Exit(1)
	}
	if err_author != nil {
		Log("error", "Author migration error.\n"+err_author.Error())
		os.Exit(1)
	}

	Log("info", "Auto migration compleate.")

}
