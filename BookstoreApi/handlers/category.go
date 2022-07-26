package handlers

import (
	"bookstore/models"
	"bookstore/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllCategories(c echo.Context) error {
	var categories = models.Categories{}
	myDB := repository.GetDB()
	categories.Categories = myDB.GetAllCategories()

	return c.JSON(http.StatusOK, categories)
}

func GetCategoryByID(c echo.Context) error {
	myDB := repository.GetDB()

	id, _ := strconv.Atoi(c.Param("id"))
	category := myDB.GetCategoryByID(id)

	return c.JSON(http.StatusOK, category)

}
func CreateCategory(c echo.Context) error {
	myDB := repository.GetDB()

	var newCategory models.Category
	if err := c.Bind(&newCategory); err != nil {
		return err
	}

	myDB.CreateCategory(0, newCategory.Name)
	newCategory = myDB.GetCategoryByName(newCategory.Name)

	return c.JSON(http.StatusOK, newCategory)
}
func PutCategory(c echo.Context) error {
	myDB := repository.GetDB()
	id, _ := strconv.Atoi(c.Param("id"))

	var newCategoryName models.Category
	if err := c.Bind(&newCategoryName); err != nil {
		return err
	}

	myDB.UpdateCategory(newCategoryName.Name, id)
	category := myDB.GetCategoryByID(id)

	return c.JSON(http.StatusOK, category)
}
func DeleteCategory(c echo.Context) error {
	myDB := repository.GetDB()

	id, _ := strconv.Atoi(c.Param("id"))
	myDB.DeleteBooksInTheCategory(id)
	deletedCategory := myDB.DeleteCategory(id)

	return c.JSON(http.StatusOK, deletedCategory)

}
