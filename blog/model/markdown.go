package model

import (
	"github.com/jinzhu/gorm"
	"html"
	"time"
)

type Markdown struct {
	Id            int    `json:"id" gorm:"primary_key"`
	Content       string `json:"content" form:"content" validate:"required"`
	IsDeleted     int    `json:"is_deleted" form:"is_deleted"`
	IsDraft       int    `json:"is_draft" form:"is_draft"`
	CreatedAt     int64  `json:"-" gorm:""`
	UpdatedAt     int64  `json:"-"`
	CreatedAtTime string `json:"created_at" gorm:"-"`
	UpdatedAtTime string `json:"updated_at" gorm:"-"`
}

func (Markdown) TableName() string {
	return "markdown"
}

func (m *Markdown) BeforeCreate(scope *gorm.Scope) {
	_ = scope.SetColumn("CreatedAt", time.Now().Unix())
}

func (m *Markdown) AfterCreate() {
	m.CreatedAtTime = time.Unix(m.CreatedAt, 0).Format(FORMAT_ISO8601)
}

func (m *Markdown) BeforeSave(scope *gorm.Scope) {
	_ = scope.SetColumn("UpdatedAt", time.Now().Unix())
	_ = scope.SetColumn("Content", html.EscapeString(m.Content))
}

func (m *Markdown) AfterSave() {
	m.UpdatedAtTime = time.Unix(m.UpdatedAt, 0).Format(FORMAT_ISO8601)
	m.Content = html.UnescapeString(m.Content)
}

func (m *Markdown) AfterFind() {
	m.Content = html.UnescapeString(m.Content)
	m.CreatedAtTime = time.Unix(m.CreatedAt, 0).Format(FORMAT_ISO8601)
	m.UpdatedAtTime = time.Unix(m.UpdatedAt, 0).Format(FORMAT_ISO8601)
}
