package main

import (
	"os"

	. "github.com/c0de-runn3r/libraryManagementSystem/controllers"
	. "github.com/c0de-runn3r/libraryManagementSystem/db"
	. "github.com/c0de-runn3r/libraryManagementSystem/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/postgres"
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

	sqlDNS := os.Getenv("SQL_DNS")

	db, err := gorm.Open(postgres.Open(sqlDNS), &gorm.Config{})
	if err != nil {
		Log("error", "Can't Connect to database.\n"+err.Error())
		os.Exit(1)
	}

	Log("info", "Connected to database.")

	if os.Getenv("MIGRATE") == "true" {
		Migrate(db)
	}

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowCredentials: true}))

	apiGroup := e.Group("/api")

	apiGroup.Use(DBMiddleware(db))

	apiGroup.POST("/register", HandleRegister)
	apiGroup.POST("/login", HandleLogin)

	usersGroup := apiGroup.Group("/users")
	booksGroup := apiGroup.Group("/books")

	AttachUsersController(usersGroup, db)
	AttachBooksController(booksGroup, db)

	Log("error", e.Start(":"+os.Getenv("LMS_BACKEND_PORT")).Error())

}
