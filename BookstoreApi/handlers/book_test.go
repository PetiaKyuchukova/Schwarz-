package handlers

import (
	"bookstore/mock"
	"bookstore/models"
	"bookstore/repository"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var expectedBooks = `{"books":[{"Book":{"ID":1,"title":"testBook","author":1,"category":1,"price":12},"Author":{"ID":0},"Category":{}}]}
`
var expectedBook = `{"ID":1,"title":"testBook","author":{"ID":0},"category":{},"price":12}
`

func TestGetAllBooks(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/books", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()
	mockDB, mock, err := mock.NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "category_id", "author_id", "price"}).AddRow(1, "testBook", 1, 1, 12)
	authorRow := sqlmock.NewRows([]string{"id", "name", "biography"}).AddRow(1, "testname", "testBio")
	categoryRow := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "myCategory")

	book := models.Book{
		Title:      "testBook",
		CategoryID: 1,
		AuthorID:   1,
		Price:      22,
	}
	author := models.Author{
		Name:      "testName",
		Biography: "testBio",
	}
	category := models.Category{
		Name: "myCategory",
	}

	mockDB.Create(&author)
	mockDB.Create(&category)
	mockDB.Create(&book)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "books" `)).
		WillReturnRows(rows)
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "authors" WHERE id = ?`)).WithArgs(1).
		WillReturnRows(authorRow)
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "categories" WHERE categories.id  = ?`)).WithArgs(1).
		WillReturnRows(categoryRow)

	repository.SetDB(mockDB)
	repository.GetDB().CreateBook(book.Title, int(book.AuthorID), int(book.CategoryID), book.Price)

	if assert.NoError(t, GetAllBooks(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedBooks, rec.Body.String())
	}

}
func TestGetBookByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/books/:id", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/books/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	res := rec.Result()

	defer res.Body.Close()
	mockDB, mock, err := mock.NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	book := models.Book{
		Title:      "testBook",
		CategoryID: 1,
		AuthorID:   1,
		Price:      22,
	}
	mockDB.Create(&book)
	rows := sqlmock.NewRows([]string{"id", "title", "category_id", "author_id", "price"}).AddRow(1, "testBook", 1, 1, 12)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "books" WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	repository.SetDB(mockDB)
	//db := repository.GetDB()
	if assert.NoError(t, GetBookByID(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedBook, rec.Body.String())
	}

}
