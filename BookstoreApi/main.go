package main

import (
	"bookstore/handlers"
	"bookstore/repository"
	"log"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=postgres password=0041129115 dbname=Bookstore port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func main() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	repository.SetDB(db)

	router := echo.New()

	router.GET("/categories", handlers.GetAllCategories)
	router.GET("/categories/:id", handlers.GetCategoryByID)
	router.POST("/categories", handlers.CreateCategory)
	router.PUT("/categories/:id", handlers.PutCategory)
	router.DELETE("/categories/:id", handlers.DeleteCategory)

	router.GET("/books", handlers.GetAllBooks)
	router.GET("/books/:id", handlers.GetBookByID)
	router.POST("/books", handlers.CreateBook)
	router.PUT("/books/:id", handlers.PutBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	router.GET("/authors", handlers.GetAllAuthors)
	router.GET("/authors/:id", handlers.GetAuthorByID)
	router.POST("/authors", handlers.CreateAuthor)
	router.PUT("/authors/:id", handlers.PutAuthor)
	router.DELETE("/authors/:id", handlers.DeleteAuthor)

	router.Logger.Fatal(router.Start(":2000"))
}
