package mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewInmemoryMock() (*gorm.DB, error) {
	mockDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	mockDB.Exec(`CREATE TABLE authors (
		id SERIAL PRIMARY KEY,
	 	name varchar NOT NULL,
		biography varchar NOT NULL
	   );`)
	mockDB.Exec(`CREATE TABLE categories (
		id SERIAL PRIMARY KEY,
	 	name varchar NOT NULL
	   );`)
	mockDB.Exec(`CREATE TABLE books (
		id SERIAL PRIMARY KEY,
	 	title varchar NOT NULL,
		author_id INT REFERENCES authors(id),
		category_id INT REFERENCES categories(id),
		price float
	   );`)

	return mockDB, err
}
func NewSQLMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return mockDB, mock, err
}
