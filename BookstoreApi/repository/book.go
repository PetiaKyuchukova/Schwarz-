package repository

import (
	"bookstore/models"
	"log"
)

func (bs *Storage) CreateBook(title string, author int, category int, price float32) {
	var exists bool
	book := models.Book{Title: title, AuthorID: uint(author), CategoryID: uint(category), Price: price}

	bs.db.Model(book).Select("count(*) > 0").Where("title = ?", title).Find(&exists)

	if exists == false {
		bs.db.Create(&book)
	} else {
		log.Print("The book already exists!")
	}

}
func (bs *Storage) GetAllBooks() []models.Book {

	books := []models.Book{}
	bs.db.Find(&books)

	return books
}
func (bs *Storage) GetBookByID(id int) models.Book {
	book := models.Book{}
	bs.db.Where("id = ?", id).Find(&book)
	return book
}
func (bs *Storage) GetBooksByAuthorId(author_id int) []models.Book {
	books := []models.Book{}

	bs.db.Where("books.author_id  = ?", author_id).Find(&books).Scan(&books)

	return books
}
func (bs *Storage) UpdateBookPrice(id int, price float32) models.Book {

	book := bs.GetBookByID(id)

	bs.db.Model(&book).Update("price", price)

	return book
}
func (bs *Storage) DeleteBook(id int) models.Book {
	book := bs.GetBookByID(id)

	bs.db.Delete(&models.Book{}, id)
	return book

}
func (bs *Storage) DeleteBooksInTheCategory(category_id int) {
	book := models.Book{}
	bs.db.Where("category_id = ?", category_id).Delete(&book)
}
func (bs *Storage) DeleteBooksOfTheAuthor(author_id int) {
	book := models.Book{}
	bs.db.Where("author_id = ?", author_id).Delete(&book)
}
