package repository

import (
	"bookstore/models"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {

	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	book := models.Book{
		Title:      "Test",
		AuthorID:   1,
		CategoryID: 1,
		Price:      12,
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "books" ("title","author_id","category_id","price") VALUES ($1,$2,$3,$4)`)).
		WillReturnRows(rows)

	mock.ExpectCommit()

	SetDB(mockDB)
	repo := GetDB()
	err = repo.CreateBook(book.Title, int(book.AuthorID), int(book.CategoryID), book.Price)
	//fmt.Println(a)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Create User: %v", err)
	}
}
func TestGetAllBooks(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "category_id", "author_id", "price"}).AddRow(1, "testBook", 1, 1, 12)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "books"`)).
		WillReturnRows(rows)

	SetDB(mockDB)
	repo := GetDB()
	authors := repo.GetAllBooks()
	assert.NotEmpty(t, authors)
	assert.Len(t, authors, 1)
}

func TestGetBookById(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "category_id", "author_id", "price"}).AddRow(1, "testBook", 1, 1, 12)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "books" WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	SetDB(mockDB)
	repo := GetDB()
	user := repo.GetBookByID(1)
	fmt.Println(user)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
}
func TestGetBooksByAuthorId(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "category_id", "author_id", "price"}).AddRow(1, "testBook", 1, 1, 12)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "books" WHERE books.author_id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	SetDB(mockDB)
	repo := GetDB()
	user := repo.GetBooksByAuthorId(1)
	fmt.Println(user)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
}
func TestGetBookByTitle(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "category_id", "author_id", "price"}).AddRow(1, "testBook", 1, 1, 12)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "books" WHERE title = $1`)).
		WithArgs("testBook").
		WillReturnRows(rows)

	SetDB(mockDB)
	repo := GetDB()
	user := repo.GetBookByTitle("testBook")
	fmt.Println(user)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
}
