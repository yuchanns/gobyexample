package service

import (
	"github.com/jinzhu/gorm"
	"github.com/yuchanns/gobyexample/model"
)

func SubQuery(DB *gorm.DB) (orderItemSub model.OrderItem) {
	DB.Where("order_id = ?", DB.Table("order").Select("id").Where("id = ?", 1).SubQuery()).Unscoped().Find(&orderItemSub)
	return
}
