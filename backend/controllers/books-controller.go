package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/c0de-runn3r/libraryManagementSystem/db/models"
	. "github.com/c0de-runn3r/libraryManagementSystem/utils"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ISBNRequest struct {
	ISBN string `json:"isbn"`
}

func AttachBooksController(g *echo.Group, db *gorm.DB) {

	Log("info", "Attaching BOOKS controller.")

	g.POST("/add-book", handleAddBook)
	g.POST("/add-book-by-isbn", handleAddBookByISBN)
	g.GET("/get-book", handleGetBook)
	g.GET("/get-all", handleGetAllBooks)

	g.POST("/take-book", handleTakeBook)
	g.POST("/return-book", handleReturnBook)
}

func handleAddBook(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	newBook := new(models.Book)
	json.NewDecoder(c.Request().Body).Decode(&newBook)

	if newBook.Title == "" {
		return c.JSON(http.StatusBadRequest, "not enough data")
	}

	for i, v := range newBook.Authors { // check if any of the authors are already in database, if so - change ID into exsisting one to prevent from duplicating
		author := new(models.Author)
		result := database.Model(&models.Author{}).Where("name = ?", v.Name).First(&author)
		if result.RowsAffected > 0 {
			newBook.Authors[i].ID = author.ID
		}
	}

	database.Create(&newBook)
	if newBook.ID == 0 {
		return c.JSON(http.StatusConflict, "book already exists")
	}
	Log("debug", "Added new book")
	return c.JSON(http.StatusOK, newBook)
}

func handleAddBookByISBN(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	isbn := new(ISBNRequest)
	json.NewDecoder(c.Request().Body).Decode(&isbn)

	if isbn.ISBN == "" {
		return c.JSON(http.StatusBadRequest, "not enough data")
	}

	newBook, err := FindBookByISBN(isbn.ISBN)
	if err != nil {
		panic("PANIC")
	}

	for i, v := range newBook.Authors { // check if any of the authors are already in database, if so - change ID into exsisting one to prevent from duplicating
		author := new(models.Author)
		result := database.Model(&models.Author{}).Where("name = ?", v.Name).First(&author)
		if result.RowsAffected > 0 {
			newBook.Authors[i].ID = author.ID
		}
	}

	database.Create(&newBook)
	if newBook.ID == 0 {
		return c.JSON(http.StatusConflict, "book already exists")
	}
	Log("debug", "Added new book by ISBN")
	return c.JSON(http.StatusOK, newBook)
}

func handleGetBook(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	var books []*models.Book

	if c.QueryParam("author") != "" && c.QueryParam("title") != "" {
		var author models.Author
		var bookIDs []int
		database.Model(&models.Author{}).Where("name = ?", c.QueryParam("author")).First(&author)
		database.Table("author_books").Where("author_id = ?", author.ID).Select("book_id").Find(&bookIDs)
		database.Model(&models.Book{}).Preload("Authors").Where("title = ?", c.QueryParam("title")).Find(&books, bookIDs)
	} else if c.QueryParam("author") != "" {
		var author models.Author
		var bookIDs []int
		database.Model(&models.Author{}).Where("name = ?", c.QueryParam("author")).First(&author)
		database.Table("author_books").Where("author_id = ?", author.ID).Select("book_id").Find(&bookIDs)
		database.Model(&models.Book{}).Preload("Authors").Find(&books, bookIDs)
	} else if c.QueryParam("title") != "" {
		database.Model(&models.Book{}).Preload("Authors").Where("title = ?", c.QueryParam("title")).Find(&books)
	}

	if len(books) == 0 {
		return c.JSON(http.StatusNoContent, "no books found")
	}

	Log("debug", "Handled get book")
	return c.JSON(http.StatusOK, books)
}

func handleGetAllBooks(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	var books []*models.Book
	database.Model(&models.Book{}).Preload("Authors").Find(&books)
	return c.JSON(http.StatusOK, books)
}
