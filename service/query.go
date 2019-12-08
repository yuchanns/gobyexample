package service

import (
	"github.com/jinzhu/gorm"
	"github.com/yuchanns/gobyexample/model"
)

func QueryPreload(DB *gorm.DB) *model.Order {
	var order model.Order
	DB.Where("id = ?", 1).Preload("OrderItems", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Unscoped().Find(&order)
	return &order
}
