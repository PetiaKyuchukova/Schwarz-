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

var expectedCategories = `{"categories":[{"id":1,"name":"myCategory"},{"id":2,"name":"myCategory"},{"id":3,"name":"myCategory"}]}
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
	mockDB, mock, err := mock.NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	categoryRow := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "myCategory").AddRow(2, "myCategory").AddRow(3, "myCategory")

	category := models.Category{
		Name: "myCategory",
	}

	mockDB.Create(&category)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "categories" `)).
		WillReturnRows(categoryRow)

	repository.SetDB(mockDB)

	if assert.NoError(t, GetAllCategories(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedCategories, rec.Body.String())
	}
}
func TestGetCategoriesByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/categories/:id", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/categories/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	res := rec.Result()

	defer res.Body.Close()
	mockDB, mock, err := mock.NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	categoryRow := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "myCategory").AddRow(2, "myCategory").AddRow(3, "myCategory")

	category := models.Category{
		Name: "myCategory",
	}
	mockDB.Create(category)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "categories" WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(categoryRow)

	repository.SetDB(mockDB)
	//db := repository.GetDB()
	if assert.NoError(t, GetCategoryByID(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedCategory, rec.Body.String())
	}

}
