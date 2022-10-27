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
		Log("error", "User migration error.\n"+err.Error())
		os.Exit(1)
	}

	// TODO добав міграцію моделей, бо толку від них немає

	Log("info", "Auto migration compleate.")

}
