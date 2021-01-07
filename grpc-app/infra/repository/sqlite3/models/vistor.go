package models

import (
	"github.com/yuchanns/gobyexample/grpc-app/common"
	"github.com/yuchanns/gobyexample/grpc-app/domain/greeter"
)

type Visitor struct {
	ID   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (Visitor) TableName() string {
	return "visitor"
}

func DoToVisitor(do *greeter.Greeter) *Visitor {
	visitor := &Visitor{}
	common.MustConvert(do, &visitor)
	return visitor
}
