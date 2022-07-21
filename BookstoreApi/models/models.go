package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Name      string
	Biography string
}
type Category struct {
	gorm.Model
	ID   int `gorm:"primaryKey"`
	Name string
}
type Book struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	Title    string
	author   int
	category int
	price    float32
}
