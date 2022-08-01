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

	bookAuthorCategoryArr := []models.BookAuthorCategory{}

	for i := 0; i < len(books); i++ {
		author = myDB.GetAuthorById(int(books[i].AuthorID))
		category = myDB.GetCategoryOfTheBook(books[i])

		bookAuthorCategoryArr = append(bookAuthorCategoryArr, models.BookAuthorCategory{Book: books[i], Category: category, Author: author})

	}

	response := models.Result{Books: bookAuthorCategoryArr}

	return c.JSON(http.StatusOK, response)

}
func GetBookByID(c echo.Context) error {
	myDB := repository.GetDB()

	id, _ := strconv.Atoi(c.Param("id"))
	book := myDB.GetBookByID(id)

	return c.JSON(http.StatusOK, myDB.SetBookDetailInfo(book))

}
func CreateBook(c echo.Context) error {
	myDB := repository.GetDB()
	var newBook models.Book
	if err := c.Bind(&newBook); err != nil {
		return err
	}
	myDB.CreateBook(0, newBook.Title, int(newBook.AuthorID), int(newBook.CategoryID), newBook.Price)
	newBook = myDB.GetBookByTitle(newBook.Title)

	return c.JSON(http.StatusOK, myDB.SetBookDetailInfo(newBook))
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
