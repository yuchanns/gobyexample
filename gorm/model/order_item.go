package model

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

// order item
type OrderItem struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Order       *Order `gorm:"foreignkey:OrderId"` // Belongs To
	OrderId     uint   `json:"order_id"`
	SID         uint   `json:"s_id"`
	UserId      uint   `json:"user_id"`
	GID         uint   `json:"g_id"`
	Name        string `json:"name" gorm:"type:varchar(50)"`
	Num         uint   `json:"num"`
	Price       uint   `json:"price"`
	Status      uint8  `json:"status" gorm:"type:tinyint(1)"`
	IsDeleted   uint   `json:"is_deleted" gorm:"type:tinyint(1)"`
	CreatedAt   int64  `json:"-"`
	UpdatedAt   int64  `json:"-"`
	DeletedAt   int64  `json:"-"`
	CreatedTime string `json:"created_time" gorm:"-"` //ignore this field
	UpdatedTime string `json:"updated_time" gorm:"-"` //ignore this field
}

func (oi *OrderItem) TableName() string {
	return "order_item"
}

func (oi *OrderItem) AfterFind() {
	oi.CreatedTime = time.Unix(oi.CreatedAt, 0).Format("2006-01-02 15:04:05")
	oi.UpdatedTime = time.Unix(oi.UpdatedAt, 0).Format("2006-01-02 15:04:05")
}

func (oi *OrderItem) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("CreatedAt", time.Now().Unix())
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	if _, ok := scope.FieldByName("Status"); ok {
		for _, status := range statusScope {
			if status == oi.Status {
				return
			}
		}
	}
	return
}

func (oi *OrderItem) BeforeUpdate(scope *gorm.Scope) (err error) {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	if _, ok := scope.FieldByName("Status"); ok {
		for _, status := range statusScope {
			if status == oi.Status {
				return
			}
		}
		return errors.New("status out of scope")
	}
	return
}
