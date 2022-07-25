package repository

import (
	"bookstore/models"
	"log"
)

func (bs *Storage) CreateBook(title string, author int, category int, price float32) {
	var exists bool
	book := models.Book{Title: title, AuthorID: uint(author), CategoryID: uint(category), Price: price}

	bs.Db.Model(book).Select("count(*) > 0").Where("title = ?", title).Find(&exists)

	if exists == false {
		bs.Db.Create(&book)
	} else {
		log.Print("The book already exists!")
	}

}
func (bs *Storage) GetAllBooks() []models.Book {

	books := []models.Book{}
	bs.Db.Find(&books)

	return books
}
func (bs *Storage) GetBookByID(id int) models.Book {
	book := models.Book{}
	bs.Db.Where("id = ?", id).Find(&book)
	return book
}
func (bs *Storage) GetBooksByAuthorId(author_id int) []models.Book {
	books := []models.Book{}

	bs.Db.Where("books.author_id  = ?", author_id).Find(&books).Scan(&books)

	return books
}
func (cs *Storage) GetBookByTitle(title string) models.Book {
	book := models.Book{}
	cs.Db.Where("title = ?", title).Find(&book)

	return book
}
func (cs *Storage) SetBookDetailInfo(book models.Book) models.BookDetailInfo {
	author := cs.GetAuthorOfTheBook(book)
	category := cs.GetCategoryOfTheBook(book)
	bookDetailInfo := models.BookDetailInfo{ID: book.ID, Title: book.Title, Author: author, Category: category, Price: book.Price}

	return bookDetailInfo
}
func (bs *Storage) UpdateBookPrice(id int, price float32) models.Book {

	book := bs.GetBookByID(id)

	bs.Db.Model(&book).Update("price", price)

	return book
}
func (bs *Storage) DeleteBook(id int) models.Book {
	book := bs.GetBookByID(id)

	bs.Db.Delete(&models.Book{}, id)
	return book

}
func (bs *Storage) DeleteBooksInTheCategory(category_id int) {
	book := models.Book{}
	bs.Db.Where("category_id = ?", category_id).Delete(&book)
}
func (bs *Storage) DeleteBooksOfTheAuthor(author_id int) {
	book := models.Book{}
	bs.Db.Where("author_id = ?", author_id).Delete(&book)
}
