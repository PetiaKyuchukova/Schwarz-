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

var expectedBooks = `{"books":[{"Book":{"ID":0,"title":"testBook","author":1,"category":1,"price":12},"Author":{"ID":0},"Category":{}}]}
`
var expectedBook = `{"ID":0,"title":"testBook","author":{"ID":0},"category":{},"price":12}
`
var expectedReqBook = `{"ID":1,"title":"testBook","author":1,"category":1,"price":12}
`

func TestGetAllBooks(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/books", http.NoBody)
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
	db.CreateBook(0, "testBook", 1, 1, 12)
	if assert.NoError(t, handlers.GetAllBooks(ctx)) {
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
	ctx.SetPath("/api/books/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	res := rec.Result()

	defer res.Body.Close()
	mockDB, err := mock.NewInmemoryMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	repository.SetDB(mockDB)
	db := repository.GetDB()
	db.CreateBook(1, "testBook", 1, 1, 12)

	if assert.NoError(t, handlers.GetBookByID(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedBook, rec.Body.String())
	}

}
func TestUpdateBook(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "http://localhost:2000/books/:id", strings.NewReader(expectedReqBook))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("books/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	res := rec.Result()
	defer res.Body.Close()

	mockDB, err := mock.NewInmemoryMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	repository.SetDB(mockDB)
	repository.GetDB().CreateBook(1, "testBook", 1, 1, 13)

	if assert.NoError(t, handlers.PutBook(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedBook, rec.Body.String())
	}

}
func TestDeleteBook(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:2000/books/:id", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("books/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	res := rec.Result()
	defer res.Body.Close()

	mockDB, err := mock.NewInmemoryMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	repository.SetDB(mockDB)
	repository.GetDB().CreateBook(1, "testBook", 1, 1, 12)

	if assert.NoError(t, handlers.DeleteBook(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedBook, rec.Body.String())
	}

}
func TestPostBook(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/books", strings.NewReader(expectedReqBook))
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
	if assert.NoError(t, handlers.CreateBook(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedBook, rec.Body.String())
	}
}
