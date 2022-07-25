package models

type Author struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `json:"name,omitempty"`
	Biography string `json:"biography,omitempty"`
	Books     []Book `json:"books,omitempty"`
}
type Authors struct {
	Authors []Author `json:"authors"`
}
