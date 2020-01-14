package squirrel

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Dao struct {
	db *sqlx.DB
}

func New() *Dao {
	db, err := sqlx.Open("mysql", "root:@/gobyexample?parseTime=true")

	if err != nil {
		panic(err)
	}

	return &Dao{db: db}
}

func (d *Dao) DB() *sqlx.DB {
	return d.db
}
