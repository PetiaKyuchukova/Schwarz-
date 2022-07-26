package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

var myDB *gorm.DB

func GetDB() *Repository {
	return &Repository{Db: myDB}
}
func SetDB(dbToSet *gorm.DB) {
	myDB = dbToSet
}
