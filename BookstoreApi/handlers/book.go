package handlers

import (
	"bookstore/models"
	"bookstore/repository"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllBooks(c echo.Context) error {

	myDB := repository.GetDB()

	books := myDB.GetAllBooks()
	author := models.Author{}
	category := models.Category{}

	response := []models.Result{}

	for i := 0; i < len(books); i++ {
		author = myDB.GetAuthorOfTheBook(books[i])
		category = myDB.GetCategoryOfTheBook(books[i])

		response = append(response, models.Result{Book: books[i], Category: category, Author: author})

	}

	return c.JSON(http.StatusOK, response)

}
