package handlers

import (
	"bookstore/handlers"
	"bookstore/mock"
	"bookstore/models"
	"bookstore/repository"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var expectedAuthors = `{"authors":[{"ID":1,"name":"testname","biography":"myAuthorTestBio"},{"ID":2,"name":"myAuthorTest2","biography":"myAuthorTestBio2"}]}
`
var expectedAuthor = `{"ID":0,"name":"myAuthorTest","biography":"myAuthorTestBio"}
`

func TestGetAllAuthors(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/authors", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()
	mockDB, mock, err := mock.NewSQLMock()
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

	repository.SetDB(mockDB)
	db := repository.GetDB()
	db.CreateAuthor(0, "AuthorName", "AuthorBiography")
	aut := db.GetAuthorByName("testName")
	fmt.Print("HERE IS MY AUT", aut)

	if assert.NoError(t, handlers.GetAllAuthors(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedAuthors, rec.Body.String())
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
	mockDB, mock, err := mock.NewSQLMock()
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
	if assert.NoError(t, handlers.GetAuthorByID(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedAuthor, rec.Body.String())
	}

}
