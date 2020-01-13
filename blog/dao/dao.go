package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	db *sql.DB
}

// get private db
func (d *Dao) DB() *sql.DB {
	return d.db
}

// get a new Dao obj
func New() *Dao {
	db, err := NewDB()

	if err != nil {
		panic(err.Error())
	}

	return &Dao{db: db}
}
