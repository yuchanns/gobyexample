package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/yuchanns/gobyexample/service"
)

var (
	DB  *gorm.DB
	err error
)

func main() {
	DB, err = gorm.Open("mysql", "root:@/gobyexample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer DB.Close()
	DB.LogMode(true)
	// insert
	//id := service.InsertGoods(DB)
	//fmt.Println("primary key is ", id)
	// query
	order := service.QueryPreload(DB)
	fmt.Printf("query result is %+v\n", order)
	fmt.Printf("the order items are %+v and %+v\n", order.OrderItems[0], order.OrderItems[1])
	// update
}
