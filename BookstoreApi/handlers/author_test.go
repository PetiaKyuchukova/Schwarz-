package handlers

import (
	"bookstore/repository"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bookstore/mock"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var expextedAuthors = `{\"authors\":[
	{
		"ID": 13,
		"name": "ede",
		"biography": "tessaxsat"
	}]}\n`

func TestGetAllAuthors(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:2000/authors", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()
	mockDB, _, err := mock.NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	repository.SetDB(mockDB)
	db := repository.GetDB()

	db.CreateAuthor("AuthorName", "AuthorBiography")
	aut := db.GetAuthorByName("AuthorName")
	fmt.Print("HERE IS MY AUT", aut)

	if assert.NoError(t, GetAllAuthors(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expextedAuthors, rec.Body.String())
	}

}
