package handlers

import (
	"bookstore/models"
	"bookstore/repository"
	"net/http"
	"strconv"

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
func PostBook(c echo.Context) error {
	myDB := repository.GetDB()
	var newBook models.Book
	if err := c.Bind(&newBook); err != nil {
		return err
	}
	myDB.CreateBook(newBook.Title, int(newBook.AuthorID), int(newBook.CategoryID), newBook.Price)
	newBook = myDB.GetBookByTitle(newBook.Title)

	return c.JSON(http.StatusOK, myDB.SetBookDetailInfo(newBook))
}
func GetBookByID(c echo.Context) error {
	myDB := repository.GetDB()

	id, _ := strconv.Atoi(c.Param("id"))
	book := myDB.GetBookByID(id)

	return c.JSON(http.StatusOK, myDB.SetBookDetailInfo(book))

}
func PutBook(c echo.Context) error {
	myDB := repository.GetDB()
	id, _ := strconv.Atoi(c.Param("id"))

	var newBookPrice models.Book
	if err := c.Bind(&newBookPrice); err != nil {
		return err
	}

	myDB.UpdateBookPrice(id, newBookPrice.Price)
	book := myDB.GetBookByID(id)

	return c.JSON(http.StatusOK, myDB.SetBookDetailInfo(book))
}
func DeleteBook(c echo.Context) error {
	myDB := repository.GetDB()

	id, _ := strconv.Atoi(c.Param("id"))

	deletedbook := myDB.DeleteBook(id)

	return c.JSON(http.StatusOK, myDB.SetBookDetailInfo(deletedbook))

}
