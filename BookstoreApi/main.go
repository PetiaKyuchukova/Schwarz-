package main

import (
	"bookstore/handlers"
	"bookstore/repository"
	"log"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Author struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Biography string
	Books     []Book
}
type Category struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Books []Book
}
type Book struct {
	ID         int `gorm:"primaryKey"`
	Title      string
	AuthorID   uint `gorm:"foringKey"`
	CategoryID uint `gorm:"foringKey"`
	Price      float32
}
type Result struct {
	Book     Book
	Author   Author
	Category Category
}

func main() {

	dsn := "host=localhost user=postgres password=0041129115 dbname=Bookstore port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	repository.SetDB(db)

	router := echo.New()

	router.GET("/categories", handlers.GetAllCategories)
	router.GET("/categories/:id", handlers.GetCategoryByID)
	router.DELETE("/categories/:id", handlers.DeleteCategory)
	router.POST("/categories", handlers.CreateCategory)

	router.GET("/books", handlers.GetAllBooks)

	router.Logger.Fatal(router.Start(":2000"))
}
