package controllers

import (
	. "lms/utils"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func AttachBooksController(group *echo.Group, db *gorm.DB) {

	Log("info", "Attaching BOOKS controller.")

}
