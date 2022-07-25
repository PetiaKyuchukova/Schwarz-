package models

type BookAuthorCategory struct {
	Book     Book
	Author   Author
	Category Category
}
type Result struct {
	Books []BookAuthorCategory `json:"books"`
}
