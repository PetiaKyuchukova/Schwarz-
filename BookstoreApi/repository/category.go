package repository

import (
	"bookstore/models"
)

func (cs *Storage) CreateCategory(name string) models.Category {
	category := models.Category{Name: name}
	cs.db.Create(category)
	return category
}
func (cs *Storage) GetAllCategories() []models.Category {
	categories := []models.Category{}

	cs.db.Find(&categories)

	for i := 0; i < len(categories); i++ {
		book := []models.Book{}

		books := cs.db.Where("books.category_id  = ?", categories[i].ID).Find(&book)
		books.Scan(&book)

		categories[i].Books = append(categories[i].Books, book...)
	}

	return categories

}
func (cs *Storage) GetCategoryByID(id int) models.Category {
	category := models.Category{}

	cs.db.Where("id = ?", id).Find(&category)
	//cs.db.Where("books.category_id  = ?", category.ID).Find(&category.Books)

	return category
}
func (cs *Storage) UpdateCategory(name string, id int) models.Category {
	category := cs.GetCategoryByID(id)
	cs.db.Model(&category).Update("name", name)

	return category
}
func (cs *Storage) DeleteCategory(id int) models.Category {
	category := cs.GetCategoryByID(id)
	cs.db.Delete(&models.Category{}, id)

	return category

}
