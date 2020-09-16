package sqlmock

import (
	"bytes"
	"crypto/rand"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"math/big"
	"time"
)

type OrderStatus uint8

const (
	OrderPending OrderStatus = iota
	OrderTransferring
	OrderSuccess
	OrderReturning
	OrderRefunded
	OrderCancelled
)

type Order struct {
	ID         int64       `json:"id" gorm:"primary_key"`
	UserID     int64       `json:"user_id"`
	GID        int64       `json:"g_id"`
	UnitPrice  int64       `json:"unit_price"`
	Count      int64       `json:"count"`
	Status     OrderStatus `json:"status"`
	TotalPrice int64       `json:"total_price"`
	CreatedAt  int64       `json:"-"`
	UpdateAt   int64       `json:"-"`
}

func (Order) TableName() string {
	return "order"
}

func (o *Order) BeforeCreate(scope *gorm.Scope) error {
	currentTime := time.Now().Unix()

	for _, column := range []string{"CreatedAt", "UpdatedAt"} {
		if err := scope.SetColumn(column, currentTime); err != nil {
			return err
		}
	}

	return nil
}

func (o *Order) MarshalJSON() ([]byte, error) {
	type Alias Order
	const layout = "2006-01-02 15:04:05"
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		Alias:     (*Alias)(o),
		CreatedAt: time.Unix(o.CreatedAt, 0).Format(layout),
		UpdatedAt: time.Unix(o.UpdateAt, 0).Format(layout),
	})
}

func QueryRows() error {
	db, mock, err := sqlmock.New()
	if err != nil {
		return err
	}

	autoGenOrder := func() func() []driver.Value {
		i := 0
		userId := 10
		good := new(big.Int).SetInt64(999)
		price := new(big.Int).SetInt64(99)
		counts := new(big.Int).SetInt64(9)
		sts := new(big.Int).SetInt64(5)
		allSts := []OrderStatus{
			OrderPending,
			OrderTransferring,
			OrderSuccess,
			OrderReturning,
			OrderRefunded,
			OrderCancelled,
		}
		currentTime := time.Now().Unix()
		return func() []driver.Value {
			i++
			gid, _ := rand.Int(rand.Reader, good)
			unitePrice, _ := rand.Int(rand.Reader, price)
			count, _ := rand.Int(rand.Reader, counts)
			totalPrice := unitePrice.Int64() * count.Int64()
			status, _ := rand.Int(rand.Reader, sts)
			return []driver.Value{
				i, userId, gid.Int64(), unitePrice.Int64(), count.Int64(),
				totalPrice, allSts[status.Int64()], currentTime, currentTime + int64(i)*price.Int64(),
			}
		}
	}()

	rows := sqlmock.NewRows([]string{
		"id", "user_id", "g_id", "unit_price", "count",
		"total_price", "status", "created_at", "update_at",
	})

	for i := 0; i < 20; i++ {
		rows.AddRow(autoGenOrder()...)
	}

	o, err := gorm.Open("mysql", db)
	if err != nil {
		return err
	}
	defer o.Close()

	o.LogMode(true)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	var results []*Order
	o.Where("id > ?", 0).Find(&results)
	jsonBytes, err := json.Marshal(results)
	if err != nil {
		return err
	}
	fmt.Println(bytes.NewBuffer(jsonBytes).String())

	return nil
}
