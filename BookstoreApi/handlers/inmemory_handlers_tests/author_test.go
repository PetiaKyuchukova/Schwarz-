package handlers

import (
	"bookstore/handlers"
	"bookstore/mock"
	"bookstore/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var expectedListOfAuthors = `{"authors":[{"ID":1,"name":"AuthorName1","biography":"AuthorBiography1"}]}
`
var expectedAuthor = `{"ID":1,"name":"AuthorName1","biography":"AuthorBiography1"}
`
var expectedUpdatedAuthor = `{"ID":1,"name":"AuthorName2","biography":"AuthorBiography2"}
`

func TestGetAllAuthors(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/authors", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()
	mockDB, err := mock.NewInmemoryMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	repository.SetDB(mockDB)
	db := repository.GetDB()
	db.CreateAuthor(1, "AuthorName1", "AuthorBiography1")

	if assert.NoError(t, handlers.GetAllAuthors(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedListOfAuthors, rec.Body.String())
	}

}
func TestGetAuthorByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/authors/:id", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/authors/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	res := rec.Result()

	defer res.Body.Close()
	mockDB, err := mock.NewInmemoryMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	repository.SetDB(mockDB)

	repository.GetDB().CreateAuthor(1, "AuthorName1", "AuthorBiography1")

	if assert.NoError(t, handlers.GetAuthorByID(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedAuthor, rec.Body.String())
	}

}
func TestUpdateAuthor(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "http://localhost:2000/authors/:id", strings.NewReader(expectedUpdatedAuthor))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("authors/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	res := rec.Result()
	defer res.Body.Close()

	mockDB, err := mock.NewInmemoryMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	repository.SetDB(mockDB)
	repository.GetDB().CreateAuthor(1, "AuthorName1", "AuthorBiography1")

	if assert.NoError(t, handlers.PutAuthor(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedUpdatedAuthor, rec.Body.String())
	}

}
func TestDeleteAuthor(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:2000/authors/:id", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("authors/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	res := rec.Result()
	defer res.Body.Close()

	mockDB, err := mock.NewInmemoryMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	repository.SetDB(mockDB)
	repository.GetDB().CreateAuthor(1, "AuthorName1", "AuthorBiography1")

	if assert.NoError(t, handlers.DeleteAuthor(ctx)) {
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
	mockDB, err := mock.NewInmemoryMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	repository.SetDB(mockDB)

	///fmt.Print("actual     ", rec.Body.String())
	if assert.NoError(t, handlers.CreateAuthor(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedAuthor, rec.Body.String())
	}
}
