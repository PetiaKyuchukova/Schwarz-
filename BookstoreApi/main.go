package main

type Author struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Biography string
	Books     []Book
}
type Category struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Books []Book
}
type Book struct {
	ID         int `gorm:"primaryKey"`
	Title      string
	AuthorID   uint `gorm:"foringKey"`
	CategoryID uint `gorm:"foringKey"`
	Price      float32
}
type Result struct {
	Book     Book
	Author   Author
	Category Category
}

func main() {
	// dsn := "host=localhost user=postgres password= dbname=Bookstore port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // //aut := models.Author{ID: 6, Name: "NewNAme", Biography: "NewBiography"}

	// fmt.Println()
	// CreateCategory(db, "categ3")

}
