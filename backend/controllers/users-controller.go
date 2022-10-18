package controllers

import (
	. "lms/utils"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func AttachUsersController(group *echo.Group, db *gorm.DB) {

	Log("info", "Attaching USERS controller.")

}
