package inmemoryrepotests

import (
	"bookstore/models"
	"bookstore/repository"
	"reflect"
	"testing"
)

//var mockDB, _ = mock.NewInmemoryMock()

func TestCreateCategory(t *testing.T) {
	expectedCategory := models.Category{ID: 1, Name: "test"}
	repository.SetDB(mockDB)
	repository.GetDB().CreateCategory(1, "test")

	dbCategory := repository.GetDB().GetCategoryByID(1)

	if reflect.DeepEqual(expectedCategory, dbCategory) == false {
		t.Error("Not expected result!", expectedCategory, dbCategory)
	}
}
func TestGetAllCategories(t *testing.T) {
	expectedCategories := []models.Category{}
	book := models.Book{
		ID:         1,
		Title:      "testBook",
		AuthorID:   1,
		CategoryID: 1,
		Price:      12,
	}
	books := []models.Book{}
	books = append(books, book)
	expectedCategory := models.Category{ID: 1, Name: "test", Books: books}

	expectedCategories = append(expectedCategories, expectedCategory)
	repository.SetDB(mockDB)
	repository.GetDB().CreateCategory(1, "test")
	repository.GetDB().CreateBook(1, "testBook", 1, 1, 12)

	result := repository.GetDB().GetAllCategories()
	if reflect.DeepEqual(expectedCategories, result) == false {
		t.Error("Not expected result!", expectedCategories, result)
	}

}
func TestGetCategoryById(t *testing.T) {
	expectedCategory := models.Category{ID: 1, Name: "test"}

	repository.SetDB(mockDB)
	repository.GetDB().CreateCategory(1, "test")

	result := repository.GetDB().GetCategoryByID(1)
	if reflect.DeepEqual(expectedCategory, result) == false {
		t.Error("Not expected result!", expectedCategory, result)
	}

}
func TestGetCategoryByName(t *testing.T) {
	expectedCategory := models.Category{ID: 1, Name: "test"}

	repository.SetDB(mockDB)
	repository.GetDB().CreateCategory(1, "test")

	result := repository.GetDB().GetCategoryByName("test")
	if reflect.DeepEqual(expectedCategory, result) == false {
		t.Error("Not expected result!", expectedCategory, result)
	}

}
func TestUpdateCategory(t *testing.T) {
	expectedCategory := models.Category{ID: 1, Name: "test"}

	repository.SetDB(mockDB)
	repository.GetDB().CreateCategory(1, "test1")

	result := repository.GetDB().UpdateCategory("test", 1)
	if reflect.DeepEqual(expectedCategory, result) == false {
		t.Error("Not expected result!", expectedCategory, result)
	}

}
func TestDeleteCategory(t *testing.T) {
	expectedCategory := models.Category{}

	repository.SetDB(mockDB)
	repository.GetDB().CreateCategory(1, "test1")
	repository.GetDB().DeleteCategory(1)
	check := repository.GetDB().GetCategoryByID(1)

	if reflect.DeepEqual(expectedCategory, check) == false {
		t.Error("Not expected result!", expectedCategory, check)
	}

}
