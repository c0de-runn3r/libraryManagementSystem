package utils

import (
	"encoding/json"
	"net/http"

	"github.com/c0de-runn3r/libraryManagementSystem/db/models"
)

const baseURLGoogleAPI = "https://www.googleapis.com/books/v1/volumes?q=isbn:"

type GoogleAPIBookResponce struct {
	Items []struct {
		VolumeInfo struct {
			Title               string   `json:"title"`
			Authors             []string `json:"authors"`
			Publisher           string   `json:"publisher"`
			PublishedDate       string   `json:"publishedDate"`
			Description         string   `json:"description"`
			IndustryIdentifiers []struct {
				Type       string `json:"type"`
				Identifier string `json:"identifier"`
			} `json:"industryIdentifiers"`
			PageCount  int      `json:"pageCount"`
			PrintType  string   `json:"printType"`
			Categories []string `json:"categories"`
			Language   string   `json:"language"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

func FindBookByISBN(isbn string) (models.Book, error) {
	var book models.Book

	var respGoogleAPI GoogleAPIBookResponce

	Log("debug", "making request to Google Books API")
	resp, err := http.Get(baseURLGoogleAPI + isbn)
	if err != nil {
		return book, err
	}

	json.NewDecoder(resp.Body).Decode(&respGoogleAPI)

	book = models.Book{
		Title:     respGoogleAPI.Items[0].VolumeInfo.Title,
		Authors:   []*models.Author{&models.Author{Name: respGoogleAPI.Items[0].VolumeInfo.Authors[0]}}, // TODO takes only 1 author - fix that
		Publisher: respGoogleAPI.Items[0].VolumeInfo.Publisher,
		Year:      respGoogleAPI.Items[0].VolumeInfo.PublishedDate,
		ISBN:      isbn,
		PageCount: respGoogleAPI.Items[0].VolumeInfo.PageCount,
	}

	return book, nil
}
