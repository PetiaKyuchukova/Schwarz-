package repository

import (
	"bookstore/models"
	"fmt"
)

func (as *Repository) CreateAuthor(name string, biography string) models.Author {
	//var exists bool
	author := models.Author{Name: name, Biography: biography}

	//as.Db.Model(author).Select("count(*) > 0").Where("name = ?", name).Find(&exists)

	//if exists == false {
	as.Db.Create(&author)
	//} else {
	//	log.Print("The author already exists!")
	//}

	return author
}
func (as *Repository) GetAllAuthors() []models.Author {
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
func (as *Repository) GetAuthorById(id int) models.Author {
	author := models.Author{}
	as.Db.Where("id = ?", id).Find(&author)

	return author
}
func (cs *Repository) GetAuthorByName(name string) models.Author {
	author := models.Author{}
	cs.Db.Where("name = ?", name).Find(&author)

	return author
}

func (as *Repository) UpdateAuthor(id int, name string, bio string) error {

	err := as.Db.Model(&models.Author{}).Where("id", id).Update("name", name).Error
	if err != nil {
		fmt.Print("err1")
		return err

	}
	err = as.Db.Model(&models.Author{}).Where("id", id).Update("biography", bio).Error
	if err != nil {
		fmt.Print("err2")

		return err
	}
	fmt.Print("nil")

	return nil
}
func (as *Repository) DeleteAuthor(id int) error {
	err := as.Db.Where("1 = 1").Delete(&models.Author{}).Error
	if err != nil {
		fmt.Print("err1")
		return err

	}
	return nil
}
