package repository

import (
	"bookstore/models"
	"log"
)

func (cs *Repository) CreateCategory(name string) models.Category {

	category := models.Category{Name: name}
	var exists bool
	cs.Db.Model(category).Select("count(*) > 0").Where("name = ?", name).Find(&exists)

	if exists == false {
		cs.Db.Create(&category)
	} else {
		log.Print("Category already exists!")
	}

	return category
}
func (cs *Repository) GetAllCategories() []models.Category {
	categories := []models.Category{}
	cs.Db.Find(&categories)

	for i := 0; i < len(categories); i++ {
		book := []models.Book{}

		books := cs.Db.Where("books.category_id  = ?", categories[i].ID).Find(&book)
		books.Scan(&book)

		categories[i].Books = append(categories[i].Books, book...)
	}

	return categories

}
func (cs *Repository) GetCategoryByID(id int) models.Category {
	category := models.Category{}
	cs.Db.Where("id = ?", id).Find(&category)

	return category
}
func (cs *Repository) GetCategoryByName(name string) models.Category {
	category := models.Category{}
	cs.Db.Where("name = ?", name).Find(&category)

	return category
}
func (cs *Repository) GetCategoryOfTheBook(book models.Book) models.Category {
	category := models.Category{}

	cs.Db.Where("categories.id  = ?", book.CategoryID).Find(&category).Scan(&category)

	return category
}
func (cs *Repository) UpdateCategory(name string, id int) models.Category {
	category := cs.GetCategoryByID(id)
	cs.Db.Model(&category).Update("name", name)

	return category
}
func (cs *Repository) DeleteCategory(id int) models.Category {
	category := cs.GetCategoryByID(id)
	cs.Db.Delete(&models.Category{}, id)

	return category

}
