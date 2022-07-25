package repository

import (
	"gorm.io/gorm"
)

type Storage struct {
	Db *gorm.DB
}

var myDB *gorm.DB

func GetDB() *Storage {
	return &Storage{Db: myDB}
}
func SetDB(dbToSet *gorm.DB) {
	myDB = dbToSet
}
