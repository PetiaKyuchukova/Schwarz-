package models

type Category struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Books []Book
}
