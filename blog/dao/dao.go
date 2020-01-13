package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Dao struct {
	sqlxdb *sqlx.DB
}

// get private db
func (d *Dao) SqlxDB() *sqlx.DB {
	return d.sqlxdb
}

// get a new Dao obj
func New() *Dao {
	db, err := NewDB()

	if err != nil {
		panic(err.Error())
	}

	return &Dao{sqlxdb: sqlx.NewDb(db, "mysql")}
}
