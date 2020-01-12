package boiler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dao *Dao

func init() {
	dao = New()
}

func TestCreateOne(t *testing.T) {
	order, err := CreateOne(dao.DB)

	assert.Nil(t, err, "failed to create one")

	fmt.Printf("create one row returns as: %+v", order)
}

func TestQueryOne(t *testing.T) {
	order, err := QueryOne(dao.DB)

	assert.Nil(t, err, "failed to query one")

	fmt.Printf("query one row results to: %+v", order)
}

func TestUpdateOne(t *testing.T) {
	rowsAff, err := UpdateOne(dao.DB)

	assert.Nil(t, err, "failed to update one")

	fmt.Println("update one affect rows:", rowsAff)
}
