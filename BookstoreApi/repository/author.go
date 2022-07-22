package repository

import (
	"bookstore/models"
	"log"
)

func (as *Storage) CreateAuthor(name string, biography string) models.Author {
	var exists bool
	author := models.Author{Name: name, Biography: biography}

	as.db.Model(author).Select("count(*) > 0").Where("name = ?", name).Find(&exists)

	if exists == false {
		as.db.Create(&author)
	} else {
		log.Print("The author already exists!")
	}
	as.db.Create(&author)

	return author
}
func (as *Storage) GetAllAuthors() []models.Author {
	authors := []models.Author{}
	as.db.Find(&authors)

	for i := 0; i < len(authors); i++ {
		book := []models.Book{}

		books := as.db.Where("books.author_id  = ?", authors[i].ID).Find(&book)
		books.Scan(&book)

		authors[i].Books = append(authors[i].Books, book...)
	}

	return authors
}
func (as *Storage) GetAuthorById(id int) models.Author {
	author := models.Author{}

	as.db.Where("id = ?", id).Find(&author)

	return author
}
func (as *Storage) UpdateAuthor(updatedAuthor models.Author) models.Author {
	author := as.GetAuthorById(updatedAuthor.ID)

	as.db.Model(&author).Update("name", updatedAuthor.Name)
	as.db.Model(&author).Update("biography", updatedAuthor.Biography)

	return author
}
func (as *Storage) DeleteAuthor(id int) models.Author {
	author := as.GetAuthorById(id)
	as.db.Delete(&models.Author{}, id)

	return author
}
func (as *Storage) GetAuthorOfTheBook(book models.Book) models.Author {
	author := models.Author{}
	as.db.Where("authors.id  = ?", book.AuthorID).Find(&author).Scan(&author)
	return author
}
