package service

import (
	"github.com/jinzhu/gorm"
	"github.com/yuchanns/gobyexample/model"
)

func UpdateAutoComplete(order *model.Order, DB *gorm.DB) (err error) {
	//DB.Set("gorm:save_associations", false).Model(&order).Unscoped().Update("Status", model.OrderPayed) // 不支持字段为int类型的UpdatedAt自动更新
	//DB.Model(&order).Updates(model.Order{Status: model.OrderTransporting, IsDeleted: 0}) // 不支持字段为int类型的UpdatedAt自动更新
	//DB.Model(&order).Unscoped().UpdateColumns(model.Order{Status: model.OrderTransporting, IsDeleted: 0})
	// 推荐写法：支持字段为int类型的UpdatedAt自动更新
	order.Status = model.OrderPayed
	err = DB.Unscoped().Select("Status").Save(&order).Error
	return
}
