package db

import (
	. "github.com/c0de-runn3r/libraryManagementSystem/db/models"
	. "github.com/c0de-runn3r/libraryManagementSystem/utils"

	"os"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	Log("info", "Auto migration starts...")

	err := db.AutoMigrate(&User{})
	if err != nil {
		Log("error", "User migration error.\n"+err.Error())
		os.Exit(1)
	}
	err = db.AutoMigrate(&Book{})
	if err != nil {
		Log("error", "Book migration error.\n"+err.Error())
		os.Exit(1)
	}
	err = db.AutoMigrate(&Author{})
	if err != nil {
		Log("error", "Author migration error.\n"+err.Error())
		os.Exit(1)
	}
	err = db.AutoMigrate(&BookLog{})
	if err != nil {
		Log("error", "User migration error.\n"+err.Error())
		os.Exit(1)
	}

	Log("info", "Auto migration compleate.")

}
