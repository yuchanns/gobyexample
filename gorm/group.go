package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func Group(DB *gorm.DB) []map[string]interface{} {
	rows, _ := DB.Table("order_item").Select("order_id").Group("order_id").Rows()
	defer rows.Close()
	m := make(map[string]interface{})
	var orderId string
	var list []map[string]interface{}
	for rows.Next() {
		rows.Scan(&orderId)
		m["order_id"] = orderId
		list = append(list, m)
		fmt.Printf("row is %+v\n", m)
	}
	return list
}
