package ginkgo

import "encoding/json"

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (b *Book) CategoryByLength() string {
	if b.Pages > 300 {
		return "NOVEL"
	}
	return "SHORT STORY"
}

func NewBookFromJSON(text string) *Book {
	var book Book
	_ = json.Unmarshal([]byte(text), &book)
	return &book
}
