package model

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

const (
	OrderPending      uint8 = 0
	OrderPayed        uint8 = 1
	OrderTransporting uint8 = 2
	OrderReceived     uint8 = 3
	OrderConfirmed    uint8 = 4
	OrderRefunding    uint8 = 5
	OrderRefunded     uint8 = 6
	OrderCancel       uint8 = 7
)

var statusScope = []uint8{
	OrderPending,
	OrderPayed,
	OrderTransporting,
	OrderReceived,
	OrderConfirmed,
	OrderRefunding,
	OrderRefunded,
	OrderCancel,
}

// main order
type Order struct {
	ID          uint         `json:"id" gorm:"primary_key"`
	OrderNo     string       `json:"order_no" gorm:"type:varchar(32)"`
	UserId      uint         `json:"user_id"`
	TotalPrice  uint         `json:"total_price"`
	Postage     uint         `json:"postage"`
	Status      uint8        `json:"status" gorm:"type:tinyint(1)"`
	IsDeleted   uint         `json:"is_deleted" gorm:"type:tinyint(1)"`
	CreatedAt   int64        `json:"-"`
	UpdatedAt   int64        `json:"-"`
	DeletedAt   int64        `json:"-"`
	OrderItems  []*OrderItem `gorm:"foreignkey:OrderId"`    // Has Many
	CreatedTime string       `json:"created_time" gorm:"-"` //ignore this field
	UpdatedTime string       `json:"updated_time" gorm:"-"` //ignore this field
}

func (o *Order) TableName() string {
	return "order"
}

func (o *Order) AfterFind() {
	o.CreatedTime = time.Unix(o.CreatedAt, 0).Format("2006-01-02 15:04:05")
	o.UpdatedTime = time.Unix(o.UpdatedAt, 0).Format("2006-01-02 15:04:05")
}

func (o *Order) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("CreatedAt", time.Now().Unix())
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	if _, ok := scope.FieldByName("Status"); ok {
		for _, status := range statusScope {
			if status == o.Status {
				return
			}
		}
		return errors.New("status out of scope")
	}
	return
}

func (o *Order) BeforeUpdate(scope gorm.Scope) (err error) {
	scope.Set("UpdatedAt", time.Now().Unix())
	if _, ok := scope.FieldByName("Status"); ok {
		for _, status := range statusScope {
			if status == o.Status {
				return
			}
		}
		return errors.New("status out of scope")
	}
	return
}
