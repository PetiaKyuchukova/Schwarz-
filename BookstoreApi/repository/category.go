package repository

import (
	"bookstore/models"

	"github.com/jinzhu/gorm"
)

type CategoryStore struct {
	db *gorm.DB
}

func NewCategoryStore(db *gorm.DB) *CategoryStore {
	return &CategoryStore{db: db}
}

func (cs *CategoryStore) CreateCategory(name string) {
	cs.db.Create(models.Category{Name: name})
}
func (cs *CategoryStore) GetAllCategories() []models.Category {
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
func (cs *CategoryStore) GetCategoryByID(id int) models.Category {
	category := models.Category{}

	cs.db.Where("id = ?", id).Find(&category)
	//cs.db.Where("books.category_id  = ?", category.ID).Find(&category.Books)

	return category
}
func (cs *CategoryStore) UpdateCategory(name string, id int) models.Category {
	category := cs.GetCategoryByID(id)
	cs.db.Model(&category).Update("name", name)

	return category
}
func (cs *CategoryStore) DeleteCategory(id int) models.Category {
	category := models.Category{}

	cs.db.Where("id = ?", id).Find(&category)
	cs.db.Delete(&models.Category{}, id)

	return category

}
