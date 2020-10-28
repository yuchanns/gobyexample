package domain

import (
	"math/rand"
	"time"
)

type News struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	AuthorID  int       `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewNews(authorID int, title, content string) *News {
	do := &News{
		ID:       rand.Int(),
		AuthorID: authorID,
		Title:    title,
		Content:  content,
	}
	return do
}
