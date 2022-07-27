package handlers

import (
	"bookstore/models"
	"bookstore/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllAuthors(c echo.Context) error {
	var authors = models.Authors{}
	myDB := repository.GetDB()
	authors.Authors = myDB.GetAllAuthors()

	return c.JSON(http.StatusOK, authors)
}
func GetAuthorByID(c echo.Context) error {
	myDB := repository.GetDB()

	id, _ := strconv.Atoi(c.Param("id"))
	author := myDB.GetAuthorById(id)
	books := myDB.GetBooksByAuthorId(id)
	author.Books = books

	return c.JSON(http.StatusOK, author)

}
func CreateAuthor(c echo.Context) error {
	myDB := repository.GetDB()

	var newAuthor models.Author
	if err := c.Bind(&newAuthor); err != nil {
		return err
	}

	myDB.CreateAuthor(newAuthor.Name, newAuthor.Biography)
	newAuthor = myDB.GetAuthorByName(newAuthor.Name)
	return c.JSON(http.StatusOK, newAuthor)

}
func PutAuthor(c echo.Context) error {
	myDB := repository.GetDB()
	id, _ := strconv.Atoi(c.Param("id"))

	var updateAuthor models.Author
	if err := c.Bind(&updateAuthor); err != nil {
		return err
	}

	err := myDB.UpdateAuthor(id, updateAuthor.Name, updateAuthor.Biography)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Please provide valid author id")
	}
	author := myDB.GetAuthorById(id)

	return c.JSON(http.StatusOK, author)
}
func DeleteAuthor(c echo.Context) error {
	myDB := repository.GetDB()

	id, _ := strconv.Atoi(c.Param("id"))
	myDB.DeleteBooksOfTheAuthor(id)
	deletedAuthor := myDB.DeleteAuthor(id)

	return c.JSON(http.StatusOK, deletedAuthor)

}
