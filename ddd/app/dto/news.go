package dto

import (
	"time"
)

type NewsDTO struct {
	ID        int       `form:"id"`
	Title     string    `form:"title" validate:"required"`
	AuthorID  int       `form:"author_id" validate:"required"`
	Content   string    `form:"content" validate:"required"`
	CreatedAt time.Time `form:"created_at"`
	UpdatedAt time.Time `form:"updated_at"`
	DeletedAt time.Time `form:"deleted_at"`
}

func (dto *NewsDTO) Validate() error {
	return validate(dto)
}
