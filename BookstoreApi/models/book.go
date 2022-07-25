package models

type Book struct {
	ID         int     `gorm:"primaryKey"`
	Title      string  `json:"title"`
	AuthorID   uint    `gorm:"foringKey" json:"author"`
	CategoryID uint    `gorm:"foringKey" json:"category"`
	Price      float32 `json:"price"`
}
type BookDetailInfo struct {
	ID       int      `gorm:"primaryKey"`
	Title    string   `json:"title"`
	Author   Author   `json:"author"`
	Category Category `json:"category"`
	Price    float32  `json:"price"`
}
type Books struct {
	Books []BookDetailInfo `json:"books"`
}
