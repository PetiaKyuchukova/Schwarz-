package handlers

import (
	"bookstore/mock"
	"bookstore/models"
	"bookstore/repository"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var expextedAuthors = `{"authors":[{"ID":1,"name":"testname","biography":"myAuthorTestBio"},{"ID":2,"name":"myAuthorTest2","biography":"myAuthorTestBio2"}]}
`
var expectedAuthor = `{"ID":1,"name":"myAuthorTest","biography":"myAuthorTestBio"}
`

func TestGetAllAuthors(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/authors", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()
	mockDB, mock, err := mock.NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	rows := sqlmock.NewRows([]string{"id", "name", "biography"}).AddRow(1, "testname", "myAuthorTestBio").AddRow(2, "myAuthorTest2", "myAuthorTestBio2")
	nameRow := sqlmock.NewRows([]string{"id", "name", "biography"}).AddRow(1, "testname", "myAuthorTestBio")
	//books := sqlmock.NewRows([]string{"id", "title", "author_id", "category_id", "price"}).AddRow(1, "bookName", 1, 1, 22).AddRow(2, "secondBookName", 2, 2, 33)

	author := models.Author{
		Name:      "testName",
		Biography: "testBio",
	}
	mockDB.Create(&author)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "authors" WHERE name = $1`)).
		WithArgs("testName").WillReturnRows(nameRow)
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "authors" `)).
		WillReturnRows(rows)
	// mock.ExpectQuery(regexp.QuoteMeta(
	// 	`SELECT * FROM "books" WHERE author_id = $1`)).WithArgs(1).
	// 	WillReturnRows(books)

	repository.SetDB(mockDB)
	db := repository.GetDB()
	db.CreateAuthor("AuthorName", "AuthorBiography")
	aut := db.GetAuthorByName("testName")
	fmt.Print("HERE IS MY AUT", aut)

	if assert.NoError(t, GetAllAuthors(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expextedAuthors, rec.Body.String())
	}

}
func TestGetAuthorByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/authors/:id", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/lists/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	res := rec.Result()

	defer res.Body.Close()
	mockDB, mock, err := mock.NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "biography"}).
		AddRow(1, "myAuthorTest", "myAuthorTestBio").AddRow(2, "myAuthorTest", "myAuthorTestBio")

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "authors" WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	repository.SetDB(mockDB)
	//db := repository.GetDB()
	if assert.NoError(t, GetAuthorByID(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedAuthor, rec.Body.String())
	}

}
func TestPostAuthor(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/authors", strings.NewReader(expectedAuthor))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()
	mockDB, mock, err := mock.NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	author := models.Author{
		Name:      "myAuthorTest",
		Biography: "myAuthorTestBio",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "biography"}).AddRow(1, "myAuthorTest", "myAuthorTestBio")
	//row := sqlmock.NewRows([]string{"id"}).AddRow(1)

	//mock.ExpectBegin()
	mockDB.Begin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "authors" ("name","biography") VALUES ($1,$2)`))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "authors" WHERE name = $1`)).
		WithArgs("myAuthorTest").WillReturnRows(rows)
	mock.ExpectCommit()

	repository.SetDB(mockDB)
	db := repository.GetDB()
	db.CreateAuthor(author.Name, author.Biography)

	if assert.NoError(t, CreateAuthor(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedAuthor, rec.Body.String())
	}
}
