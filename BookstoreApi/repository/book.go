package repository

import (
	"bookstore/models"
)

func (bs *Storage) CreateBook(title string, author int, category int, price float32) {
	bs.db.Create(models.Book{Title: title, AuthorID: uint(author), CategoryID: uint(category), Price: price})
}
func (bs *Storage) GetAllBooks() []models.Result {
	results := []models.Result{}

	books := []models.Book{}
	bs.db.Find(&books)

	for i := 0; i < len(books); i++ {
		result := models.Result{}
		result.Book = books[i]

		bs.db.Where("authors.id  = ?", result.Book.AuthorID).Find(&result.Author).Scan(&result.Author)
		bs.db.Where("categories.id  = ?", result.Book.CategoryID).Find(&result.Category).Scan(&result.Category)

		results = append(results, result)
	}

	return results
}
func (bs *Storage) GetBookByID(id int) models.Book {
	book := models.Book{}
	bs.db.Where("id = ?", id).Find(&book)
	return book
}
func (bs *Storage) UpdateBookPrice(id int, price float32) models.Result {
	result := models.Result{}

	bs.GetBookByID(id)

	bs.db.Model(&result.Book).Update("price", price)
	bs.db.Where("authors.id  = ?", result.Book.AuthorID).Find(&result.Author).Scan(&result.Author)
	bs.db.Where("categories.id  = ?", result.Book.CategoryID).Find(&result.Category).Scan(&result.Category)

	return result
}
func (bs *Storage) DeleteBook(id int) models.Book {
	book := bs.GetBookByID(id)

	bs.db.Delete(&models.Book{}, id)
	return book

}
