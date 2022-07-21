package models

type Book struct {
	ID         int `gorm:"primaryKey"`
	Title      string
	AuthorID   uint `gorm:"foringKey"`
	CategoryID uint `gorm:"foringKey"`
	Price      float32
}
