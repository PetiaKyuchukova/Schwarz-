package main

import (
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

func GetAuthorsBooksByItsName(db *gorm.DB, author Author) {
	db.Where("name = ?", "Elif Shafak").Find(&author)
	db.Find(&author.Books, &author.ID)
}

func main() {
	dsn := "host=localhost user=postgres password=0041129115 dbname=Bookstore port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	author := Author{}

	// category := Category{Name: "test"}
	// book := Book{Title: "test", AuthorID: 1, CategoryID: 1, Price: 0}

	// db.Migrator().CreateTable(&author)
	// db.Migrator().CreateTable(&category)
	// db.Migrator().CreateTable(&book)
	db.Where("id = ?", 2).Find(&author)

	//db.Create(&author)
	// db.Create(&category)
	// db.Create(&book)
	//sauthor.ID = db.Find(&author.ID, &author.ID)
	//fmt.Println(DeleteCategory(db, 3))
	//DeleteCategory(db, 2)
}
