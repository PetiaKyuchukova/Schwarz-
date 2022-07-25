package repository

import (
	"bookstore/models"
	"log"
)

func (as *Storage) CreateAuthor(name string, biography string) models.Author {
	var exists bool
	author := models.Author{Name: name, Biography: biography}

	as.Db.Model(author).Select("count(*) > 0").Where("name = ?", name).Find(&exists)

	if exists == false {
		as.Db.Create(&author)
	} else {
		log.Print("The author already exists!")
	}
	as.Db.Create(&author)

	return author
}
func (as *Storage) GetAllAuthors() []models.Author {
	authors := []models.Author{}
	as.Db.Find(&authors)

	for i := 0; i < len(authors); i++ {
		book := []models.Book{}

		books := as.Db.Where("books.author_id  = ?", authors[i].ID).Find(&book)
		books.Scan(&book)

		authors[i].Books = append(authors[i].Books, book...)
	}

	return authors
}
func (as *Storage) GetAuthorById(id int) models.Author {
	author := models.Author{}

	as.Db.Where("id = ?", id).Find(&author)

	return author
}
func (as *Storage) UpdateAuthor(updatedAuthor models.Author) models.Author {
	author := as.GetAuthorById(updatedAuthor.ID)

	as.Db.Model(&author).Update("name", updatedAuthor.Name)
	as.Db.Model(&author).Update("biography", updatedAuthor.Biography)

	return author
}
func (as *Storage) DeleteAuthor(id int) models.Author {
	author := as.GetAuthorById(id)
	as.Db.Delete(&models.Author{}, id)

	return author
}
func (as *Storage) GetAuthorOfTheBook(book models.Book) models.Author {
	author := models.Author{}
	as.Db.Where("authors.id  = ?", book.AuthorID).Find(&author).Scan(&author)
	return author
}
