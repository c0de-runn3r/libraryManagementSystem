package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/c0de-runn3r/libraryManagementSystem/db/models"
	. "github.com/c0de-runn3r/libraryManagementSystem/utils"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func AttachBooksController(g *echo.Group, db *gorm.DB) {

	Log("info", "Attaching BOOKS controller.")

	g.POST("/add-book", handleAddBook)
	g.GET("/get-book", handleGetBook)

	g.POST("/take-book", handleTakeBook)
	g.POST("/return-book", handleReturnBook)
}

func handleAddBook(c echo.Context) error {
	database := c.Get(dbContextKey).(*gorm.DB)

	user, err := GetUserByRequestContext(c, database)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}
	if user.Role != models.Librarian && user.Role != models.Manager {
		return c.JSON(http.StatusForbidden, "you have no roots to preform this request")
	}

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
