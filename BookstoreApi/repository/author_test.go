package repository

import (
	"bookstore/models"
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

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "authors" ("name","biography") VALUES ($1,$2)`)).
		WillReturnRows(rows)
	mock.ExpectCommit()

	SetDB(mockDB)
	repo := GetDB()
	a := repo.CreateAuthor(author.Name, author.Biography)
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

	SetDB(mockDB)
	repo := GetDB()
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

	rows := sqlmock.NewRows([]string{"id", "name", "biography"}).AddRow(1, "myAuthorTest", "myAuthorTestBio").AddRow(2, "myAuthorTest2", "myAuthorTestBio2")

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "authors"`)).
		WillReturnRows(rows)

	SetDB(mockDB)
	repo := GetDB()
	authors := repo.GetAllAuthors()
	assert.NotEmpty(t, authors)
	assert.Len(t, authors, 2)

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

	SetDB(mockDB)
	repo := GetDB()
	authors := repo.GetAuthorByName("myAuthorTest")
	expected := models.Author{ID: 1, Name: "myAuthorTest", Biography: "myAuthorTestBio"}
	assert.NotEmpty(t, authors)
	assert.Equal(t, expected, authors)
}
func TestUpdateAuthor(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	sqlmock.NewRows([]string{"id", "name", "biography"}).AddRow(1, "myAuthorTest", "myAuthorTestBio").AddRow(2, "myAuthorTest2", "myAuthorTestBio2")
	mock.ExpectBegin()
	mock.ExpectPrepare(
		`UPDATE "authors"
		SET "name" =$1
		WHERE "id" =$2`).ExpectExec().
		WithArgs("myAuthorTest", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	SetDB(mockDB)
	repo := GetDB()
	err = repo.UpdateAuthor(1, "myAuthorTest", "myAuthorTestBio")
	assert.NoError(t, err)
}
