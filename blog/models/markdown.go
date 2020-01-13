package models

type Markdown struct {
	ID            int    `json:"id" db:"id" form:"id"`
	Content       string `json:"content" db:"content" form:"content"`
	IsDeleted     int    `json:"is_deleted" db:"is_deleted" form:"is_deleted"`
	IsDraft       int    `json:"is_draft" db:"is_draft" form:"is_draft"`
	CreatedAt     int64  `json:"-" db:"created_at"`
	UpdatedAt     int64  `json:"-" db:"updated_at"`
	CreatedAtTime string `json:"created_at" db:"-"`
	UpdatedAtTime string `json:"updated_at" db:"-"`
}
