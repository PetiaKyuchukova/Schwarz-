package models

type Category struct {
	ID    int    `gorm:"primaryKey" json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Books []Book `json:"books,omitempty"`
}
