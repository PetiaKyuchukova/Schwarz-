package inmemoryrepotests

import (
	"bookstore/mock"
	"bookstore/models"
	"bookstore/repository"
	"reflect"
	"testing"
)

var mockDB, _ = mock.NewInmemoryMock()
var expectedListOfAuthors = `{"authors":[{"ID":1,"name":"AuthorName1","biography":"AuthorBiography1"}]}
`
var books = `{"ID":1,"title":"testBook","author":1,"category":1,"price":12}
`
var allAuthorsJSON = `[{1 test test [{1 testBook 1 1 12}]}]`
var expectedAuthor = `{1,"test","test"{}}
`

func TestCreateAuthor(t *testing.T) {
	expectedUser := models.Author{
		ID:        1,
		Name:      "test",
		Biography: "test",
	}
	repository.SetDB(mockDB)
	repository.GetDB().CreateAuthor(1, "test", "test")

	dbAuthor := repository.GetDB().GetAuthorById(1)

	if reflect.DeepEqual(expectedUser, dbAuthor) == false {
		t.Error("Not expected result!", expectedUser, dbAuthor)
	}
}
func TestGetAllAuthors(t *testing.T) {
	expectedAuthors := []models.Author{}
	author := models.Author{
		ID:        1,
		Name:      "test",
		Biography: "test",
		Books:     []models.Book{},
	}
	expectedAuthors = append(expectedAuthors, author)
	repository.SetDB(mockDB)
	repository.GetDB().CreateAuthor(1, "test", "test")
	//repository.GetDB().CreateBook(1, "testBook", 1, 1, 12)

	result := repository.GetDB().GetAllAuthors()
	if reflect.DeepEqual(expectedAuthors, result) == false {
		t.Error("Not expected result!", expectedAuthors, result)
	}

}
func TestGetAuthorById(t *testing.T) {
	expectedAuthor := models.Author{
		ID:        1,
		Name:      "test",
		Biography: "test",
	}

	repository.SetDB(mockDB)
	repository.GetDB().CreateAuthor(1, "test", "test")

	result := repository.GetDB().GetAuthorById(1)
	if reflect.DeepEqual(expectedAuthor, result) == false {
		t.Error("Not expected result!", expectedAuthor, result)
	}

}
func TestGetAuthorByName(t *testing.T) {
	expectedAuthor := models.Author{
		ID:        1,
		Name:      "test",
		Biography: "test",
	}

	repository.SetDB(mockDB)
	repository.GetDB().CreateAuthor(1, "test", "test")

	result := repository.GetDB().GetAuthorByName("test")
	if reflect.DeepEqual(expectedAuthor, result) == false {
		t.Error("Not expected result!", expectedAuthor, result)
	}

}
func TestUpdateAuthor(t *testing.T) {
	expectedAuthor := models.Author{
		ID:        1,
		Name:      "test",
		Biography: "test",
	}

	repository.SetDB(mockDB)
	repository.GetDB().CreateAuthor(1, "test1", "test1")

	err := repository.GetDB().UpdateAuthor(1, expectedAuthor.Name, expectedAuthor.Biography)
	if err != nil {
		t.Error("Result is not expected!")
	}

}
func TestDeleteAuthor(t *testing.T) {
	expectedAuthor := models.Author{}

	repository.SetDB(mockDB)
	repository.GetDB().CreateAuthor(1, "test1", "test1")
	repository.GetDB().DeleteAuthor(1)
	check := repository.GetDB().GetAuthorById(1)

	if reflect.DeepEqual(expectedAuthor, check) == false {
		t.Error("Not expected result!", expectedAuthor, check)
	}

}
