package handlers

import (
	"bookstore/models"
	"bookstore/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllAuthors(c echo.Context) error {
	myDB := repository.GetDB()
	authors := myDB.GetAllAuthors()

	return c.JSON(http.StatusOK, authors)
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
func GetAuthorByID(c echo.Context) error {
	myDB := repository.GetDB()

	id, _ := strconv.Atoi(c.Param("id"))
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
func PutAuthor(c echo.Context) error {
	myDB := repository.GetDB()
	id, _ := strconv.Atoi(c.Param("id"))

	var updateAuthor models.Author
	if err := c.Bind(&updateAuthor); err != nil {
		return err
	}

	myDB.UpdateAuthor(updateAuthor)
	author := myDB.GetAuthorById(id)

	return c.JSON(http.StatusOK, author)
}
