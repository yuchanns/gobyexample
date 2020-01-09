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

func (m *Markdown) AfterCreate(*gorm.Scope) {
	m.CreatedAtTime = time.Unix(m.CreatedAt, 0).Format("2006-01-02 15:04:05")
}

func (m *Markdown) BeforeSave(scope *gorm.Scope) {
	_ = scope.SetColumn("UpdatedAt", time.Now().Unix())
	_ = scope.SetColumn("Content", html.EscapeString(m.Content))
}

func (m *Markdown) AfterSave(*gorm.Scope) {
	m.UpdatedAtTime = time.Unix(m.UpdatedAt, 0).Format("2006-01-02 15:04:05")
	m.Content = html.UnescapeString(m.Content)
}
