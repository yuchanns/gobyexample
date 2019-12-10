package service

import (
	"github.com/jinzhu/gorm"
	"github.com/yuchanns/gobyexample/model"
)

func Transaction(order *model.Order, DB *gorm.DB) {
	tx := DB.Begin()
	InsertGoods(tx)
	tx.Unscoped().Delete(order)
	tx.Rollback()
}
