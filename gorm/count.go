package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/yuchanns/gobyexample/gorm/model"
)

func Count(DB *gorm.DB) (count int) {
	DB.Model(model.OrderItem{}).Unscoped().Where("order_id = ?", 1).Count(&count)
	return
}
