package inmemoryrepotests

import (
	"bookstore/models"
	"bookstore/repository"
	"reflect"
	"testing"
)

//var mockDB, _ = mock.NewInmemoryMock()

func TestCreateBook(t *testing.T) {
	expectedBook := models.Book{
		ID:         1,
		Title:      "testBook",
		AuthorID:   1,
		CategoryID: 1,
		Price:      12,
	}
	repository.SetDB(mockDB)
	repository.GetDB().CreateBook(1, expectedBook.Title, 1, 1, 12)

	dbBook := repository.GetDB().GetBookByID(1)

	if reflect.DeepEqual(expectedBook, dbBook) == false {
		t.Error("Not expected result!", expectedBook, dbBook)
	}
}
func TestGetAllBooks(t *testing.T) {
	expectedBooks := []models.Book{}
	book := models.Book{
		ID:         1,
		Title:      "testBook",
		AuthorID:   1,
		CategoryID: 1,
		Price:      12,
	}

	expectedBooks = append(expectedBooks, book)

	repository.SetDB(mockDB)
	repository.GetDB().CreateBook(1, "testBook", 1, 1, 12)

	result := repository.GetDB().GetAllBooks()
	if reflect.DeepEqual(expectedBooks, result) == false {
		t.Error("Not expected result!", expectedBooks, result)
	}

}
func TestGetBookById(t *testing.T) {
	expectedBook := models.Book{
		ID:         1,
		Title:      "testBook",
		AuthorID:   1,
		CategoryID: 1,
		Price:      12,
	}
	repository.SetDB(mockDB)
	repository.GetDB().CreateBook(1, "testBook", 1, 1, 12)

	result := repository.GetDB().GetBookByID(1)
	if reflect.DeepEqual(expectedBook, result) == false {
		t.Error("Not expected result!", expectedBook, result)
	}

}
func TestGetBookByTitle(t *testing.T) {
	expectedBook := models.Book{
		ID:         1,
		Title:      "testBook",
		AuthorID:   1,
		CategoryID: 1,
		Price:      12,
	}
	repository.SetDB(mockDB)
	repository.GetDB().CreateBook(1, "testBook", 1, 1, 12)

	result := repository.GetDB().GetBookByTitle("testBook")
	if reflect.DeepEqual(expectedBook, result) == false {
		t.Error("Not expected result!", expectedBook, result)
	}

}
func TestGetBookByAuthorID(t *testing.T) {
	expectedBooks := []models.Book{}
	book := models.Book{
		ID:         1,
		Title:      "testBook",
		AuthorID:   1,
		CategoryID: 1,
		Price:      12,
	}

	expectedBooks = append(expectedBooks, book)
	repository.SetDB(mockDB)
	repository.GetDB().CreateBook(1, "testBook", 1, 1, 12)

	result := repository.GetDB().GetBooksByAuthorId(1)
	if reflect.DeepEqual(expectedBooks, result) == false {
		t.Error("Not expected result!", expectedBooks, result)
	}

}
func TestUpdateBook(t *testing.T) {
	expectedBook := models.Book{
		ID:         1,
		Title:      "testBook",
		AuthorID:   1,
		CategoryID: 1,
		Price:      12,
	}
	repository.SetDB(mockDB)
	repository.GetDB().CreateBook(1, "testBook", 1, 1, 13)
	result := repository.GetDB().UpdateBookPrice(1, 12)

	if reflect.DeepEqual(expectedBook, result) == false {
		t.Error("Not expected result!", expectedBook, result)
	}

}
func TestDeleteBook(t *testing.T) {
	expectedBook := models.Book{}

	repository.SetDB(mockDB)
	repository.GetDB().CreateBook(1, "testBook", 1, 1, 13)
	repository.GetDB().DeleteBook(1)
	check := repository.GetDB().GetBookByID(1)

	if reflect.DeepEqual(expectedBook, check) == false {
		t.Error("Not expected result!", expectedBook, check)
	}

}
