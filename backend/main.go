package main

import (
	"os"

	. "lms/controllers"
	. "lms/db"
	. "lms/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		Log("error", "Can't load .env file.\n"+err.Error())
		os.Exit(1)
	}

	Log("info", "Starting LMS.")

	Log("info", "Connecting to database...")

	db, err := gorm.Open(mysql.Open(os.Getenv("SQL_DSN")), &gorm.Config{})

	if err != nil {
		Log("error", "Can't Connect to database.\n"+err.Error())
		os.Exit(1)
	}

	Log("info", "Connected to database.")

	if os.Getenv("MIGRATE") == "true" {
		Migrate(db)
	}

	e := echo.New()

	apiGroup := e.Group("/api")

	usersGroup := apiGroup.Group("/users")

	booksGroup := apiGroup.Group("/books")

	AttachUsersController(usersGroup, db)
	AttachBooksController(booksGroup, db)

	Log("error", e.Start(":"+os.Getenv("LMS_BACKEND_PORT")).Error())

}
