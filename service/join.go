package service

import "github.com/jinzhu/gorm"

type OrderJoin struct {
	ID      uint64
	Name    string
	OrderNo string
}

func Join(DB *gorm.DB) []*OrderJoin {
	var orderJoins []*OrderJoin
	DB.Table("order_item").Joins("left join `order` on order.id = order_item.order_id").Select("order_item.id, name, order_no").Find(&orderJoins)
	return orderJoins
}
