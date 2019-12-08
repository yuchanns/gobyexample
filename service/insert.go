package service

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/yuchanns/gobyexample/model"
)

var Node *snowflake.Node

func init() {
	var err error
	Node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err.Error())
	}
}

func InsertGoods(DB *gorm.DB) uint {
	var userId, sId uint = 1088, 2
	orderItems := []*model.OrderItem{
		{
			SId:    sId,
			UserId: userId,
			GId:    20,
			Name:   "FoodA",
			Num:    2,
			Price:  2000,
			Status: model.OrderPending,
		},
		{
			SId:    sId,
			UserId: userId,
			GId:    21,
			Name:   "FoodB",
			Num:    1,
			Price:  5000,
			Status: model.OrderPending,
		},
	}

	var totalPrice uint = 0
	for _, orderItem := range orderItems {
		totalPrice += orderItem.Price * orderItem.Num
	}

	order := model.Order{
		OrderNo:    Node.Generate().String(),
		SId:        2,
		UserId:     1088,
		TotalPrice: totalPrice,
		Postage:    1000,
		Status:     model.OrderPending,
		OrderItems: orderItems,
	}

	DB.Create(&order)

	fmt.Println("order items primary key is ", orderItems[0].Id, " and ", orderItems[1].Id)
	return order.Id
}
