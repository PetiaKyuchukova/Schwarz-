package repository

import (
	"bookstore/models"
	"bookstore/repository"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBCategory(t *testing.T) {

	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	category := models.Category{
		Name: "test",
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "categories" ("name") VALUES ($1`)).
		WillReturnRows(rows)

	mock.ExpectCommit()

	repository.SetDB(mockDB)
	repo := repository.GetDB()
	err = repo.CreateCategory(0, category.Name)
	//fmt.Println(a)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Create User: %v", err)
	}
}

func TestGetAllCategories(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "testBook")
	books := sqlmock.NewRows([]string{"id", "title", "category_id", "author_id", "price"}).AddRow(1, "testBook", 1, 1, 12)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "categories"`)).
		WillReturnRows(rows)
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "books" WHERE books.category_id  = $1`)).
		WillReturnRows(books)

	repository.SetDB(mockDB)
	repo := repository.GetDB()
	authors := repo.GetAllCategories()
	assert.NotEmpty(t, authors)
	assert.Len(t, authors, 1)
}

func TestGetById(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "categories" WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	repository.SetDB(mockDB)
	repo := repository.GetDB()
	category := repo.GetCategoryByID(1)
	fmt.Println(category)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetCategoryOfTheBook(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test")

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "categories" WHERE name = $1`)).
		WithArgs("test").
		WillReturnRows(rows)

	repository.SetDB(mockDB)
	repo := repository.GetDB()
	category := repo.GetCategoryByName("test")
	fmt.Println(category)
	if err != nil {
		t.Fatal(err)
	}

}
