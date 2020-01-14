package squirrel

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	dao := New()

	order, err := Get(dao, 1)

	assert.Nil(t, err)

	fmt.Printf("query get a result:%+v\n", order)
}

func TestInsert(t *testing.T) {
	dao := New()

	now := time.Now().Unix()

	order := &Order{
		OrderNo:    strconv.FormatInt(time.Now().UnixNano(), 10),
		UserId:     2,
		TotalPrice: 10000,
		Postage:    20,
		Status:     1,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	err := Insert(dao, order)

	assert.Nil(t, err)
	fmt.Printf("insert result:%+v\n", order)
}
