package gorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open("mysql", "root:@/gobyexample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	DB.LogMode(true)
}

func TestInsertGoods(t *testing.T) {
	id := InsertGoods(DB)
	fmt.Println("order primary key is ", id)
}

func TestQueryPreload(t *testing.T) {
	order := QueryPreload(DB)
	fmt.Printf("query result is %+v\n", order)
	fmt.Printf("the order items are %+v and %+v\n", order.OrderItems[0], order.OrderItems[1])
}

func TestUpdateAutoComplete(t *testing.T) {
	order := QueryPreload(DB)
	err = UpdateAutoComplete(order, DB)
	if err != nil {
		fmt.Println("update failed")
	}
}

func TestTransaction(t *testing.T) {
	order := QueryPreload(DB)
	Transaction(order, DB)
}

func TestJoin(t *testing.T) {
	orderJoins := Join(DB)
	fmt.Printf("join result is %+v\n", orderJoins[0])
}

func TestGroup(t *testing.T) {
	list := Group(DB)
	fmt.Printf("result is %+v\n", list)
}

func TestCount(t *testing.T) {
	count := Count(DB)
	fmt.Println("count is", count)
}

func TestSubQuery(t *testing.T) {
	orderItemSub := SubQuery(DB)
	fmt.Printf("subquery result is %+v\n", orderItemSub)
}
