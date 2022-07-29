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

var expectedCategories = `{"categories":[{"id":1,"name":"myCategory"}]}
`
var expectedCategory = `{"id":1,"name":"myCategory"}
`

func TestGetAllCategories(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/categories", http.NoBody)
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
	db.CreateCategory(1, "myCategory")

	if assert.NoError(t, handlers.GetAllCategories(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedCategories, rec.Body.String())
	}

}
func TestGetCategoryByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/category/:id", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/category/:id")
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
	db.CreateCategory(1, "myCategory")

	if assert.NoError(t, handlers.GetCategoryByID(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedCategory, rec.Body.String())
	}

}
func TestUpdateCategory(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "http://localhost:2000/category/:id", strings.NewReader(expectedCategory))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("category/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	res := rec.Result()
	defer res.Body.Close()

	mockDB, err := mock.NewInmemoryMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	repository.SetDB(mockDB)
	repository.GetDB().CreateCategory(1, "myCategory1")

	if assert.NoError(t, handlers.PutCategory(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedCategory, rec.Body.String())
	}

}
func TestDeleteCategory(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:2000/category/:id", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("category/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	res := rec.Result()
	defer res.Body.Close()

	mockDB, err := mock.NewInmemoryMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	repository.SetDB(mockDB)
	repository.GetDB().CreateCategory(1, "myCategory")

	if assert.NoError(t, handlers.DeleteCategory(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedCategory, rec.Body.String())
	}

}
func TestPostCategory(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/category", strings.NewReader(expectedCategory))
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
	if assert.NoError(t, handlers.CreateCategory(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedCategory, rec.Body.String())
	}
}
