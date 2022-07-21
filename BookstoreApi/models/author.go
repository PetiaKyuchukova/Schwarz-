package models

type Author struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Biography string
	Books     []Book
}
