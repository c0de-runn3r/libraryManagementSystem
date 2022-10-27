package db

import (
	. "lms/db/models"
	. "lms/utils"

	"os"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	Log("info", "Auto migration starts...")

	err := db.AutoMigrate(&User{})

	if err != nil {
		Log("error", "Migration error.\n"+err.Error())
		os.Exit(1)
	}

	Log("info", "Auto migration compleate.")

}
