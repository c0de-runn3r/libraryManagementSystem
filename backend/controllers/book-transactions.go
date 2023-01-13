package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/c0de-runn3r/libraryManagementSystem/db/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type BookRequest struct {
	BookID int `json:"book_id"`
}

func handleTakeBook(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	user, err := GetUserByRequestContext(c, database)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	var bookReq BookRequest
	json.NewDecoder(c.Request().Body).Decode(&bookReq)
	var book models.Book
	database.Where("id = ?", bookReq.BookID).First(&book)

	if !book.Avialable {
		return c.JSON(http.StatusBadRequest, "book is already taken")
	}
	bookLog := models.BookLog{
		BookID: uint(bookReq.BookID),
		UserID: user.ID,
	}
	database.Create(&bookLog)
	book.Avialable = false
	database.Save(&book)
	resp := fmt.Sprintf("book: %s was succesfully taken by user: %s %s", book.Title, user.Name, user.Surname)
	return c.JSON(http.StatusOK, resp)
}

func handleReturnBook(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	// user, err := GetUserByRequestContext(c, database)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, err.Error())
	// }
	var bookData BookRequest
	json.NewDecoder(c.Request().Body).Decode(&bookData)

	var bookLog models.BookLog
	database.Where("book_id = ?", bookData.BookID).Preload("User").Preload("Book").Last(&bookLog)
	if bookLog.Returned {
		return c.JSON(http.StatusBadRequest, "book is already returned")
	}
	bookLog.Returned = true
	database.Save(&bookLog)
	database.First(&models.Book{}, bookData.BookID).Update("avialable", true)
	resp := fmt.Sprintf("book: %s, user: %s %s was returned", bookLog.Book.Title, bookLog.User.Name, bookLog.User.Surname)
	return c.JSON(http.StatusOK, resp)
}
