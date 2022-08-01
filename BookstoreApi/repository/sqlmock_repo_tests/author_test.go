package repository

import (
	"bookstore/models"
	"bookstore/repository"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return mockDB, mock, err
}

func TestCreateAuthor(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	author := models.Author{
		Name:      "testName",
		Biography: "testBio",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "biography"}).AddRow(1, author.Name, author.Biography)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "authors" ("name","biography") VALUES ($1,$2)`)).
		WillReturnRows(rows)
	mock.ExpectCommit()

	repository.SetDB(mockDB)
	repo := repository.GetDB()
	a := repo.CreateAuthor(0, author.Name, author.Biography)

	fmt.Println(a)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Create User: %v", err)
	}
}
func TestGetAuthorById(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "biography"}).
		AddRow(1, "myAuthorTest", "myAuthorTestBio")

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "authors" WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	repository.SetDB(mockDB)
	repo := repository.GetDB()
	user := repo.GetAuthorById(1)
	fmt.Println(user)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
}
func TestGetAllAuthors(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "biography"}).AddRow(1, "myAuthorTest", "myAuthorTestBio")

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "authors"`)).
		WillReturnRows(rows)
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "books" WHERE books.author_id  = $1`)).WithArgs(1).
		WillReturnRows(rows)

	repository.SetDB(mockDB)
	repo := repository.GetDB()
	authors := repo.GetAllAuthors()
	assert.NotEmpty(t, authors)
	assert.Len(t, authors, 1)

}
func TestGetAuthorByName(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "biography"}).AddRow(1, "myAuthorTest", "myAuthorTestBio").AddRow(2, "myAuthorTest2", "myAuthorTestBio2")

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "authors" WHERE name = $1`)).
		WithArgs("myAuthorTest").
		WillReturnRows(rows)

	repository.SetDB(mockDB)
	repo := repository.GetDB()
	authors := repo.GetAuthorByName("myAuthorTest")
	expected := models.Author{ID: 1, Name: "myAuthorTest", Biography: "myAuthorTestBio"}
	assert.NotEmpty(t, authors)
	assert.Equal(t, expected, authors)
}
