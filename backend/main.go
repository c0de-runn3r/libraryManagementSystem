package main

import (
	"os"

	. "lms/controllers"
	. "lms/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		Log("error", "Can't load .env file.")
		os.Exit(1)
	}

	Log("info", "Starting LMS.")

	e := echo.New()

	apiGroup := e.Group("/api")

	usersGroup := apiGroup.Group("/users")

	booksGroup := apiGroup.Group("/books")

	AttachUsersController(usersGroup)
	AttachBooksController(booksGroup)

	Log("error", e.Start(":"+os.Getenv("LMS_BACKEND_PORT")).Error())

}
